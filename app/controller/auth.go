package controller

import (
	"fmt"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/kevinoctavian/evodka_backend/app/model"
	"github.com/kevinoctavian/evodka_backend/app/repository"
	"github.com/kevinoctavian/evodka_backend/pkg/config"
	"github.com/kevinoctavian/evodka_backend/pkg/validator"
	"github.com/kevinoctavian/evodka_backend/platform/database"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	user := &model.CreateUser{Role: "User"} // Default role is "user"
	// Parse the request body into the user model
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Invalid request body",
		})
	}

	validate := validator.NewValidator()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":   "Validation failed",
			"error": err.Error(),
		})
	}

	// check exists user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	userRepo := repository.NewUserRepo(database.GetDB())

	var exists bool

	exists, err = userRepo.Exists(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Failed to check if user exists",
		})
	}

	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"msg": fmt.Sprintf("User with email %s already exists try with another email", user.Email),
		})
	}

	user.Password = string(hashedPassword)
	err = userRepo.Create(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Failed to create user " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "User registered successfully",
	})
}

func Login(c *fiber.Ctx) error {
	user := model.LoginUser{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Invalid request body",
		})
	}

	validate := validator.NewValidator()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":   "Validation failed",
			"error": err.Error(),
		})
	}
	// check if user exists
	userRepo := repository.NewUserRepo(database.GetDB())
	existingUser, err := userRepo.FindByEmail(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Failed to find user",
		})
	}
	// check if user exists and password is correct
	// if user not found or password is incorrect, return unauthorized
	if existingUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg": "Invalid email or password",
		})
	}
	// compare password
	// if password is incorrect, return unauthorized
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg": "Invalid email or password",
		})
	}

	// generate access token
	accessToken, err := generateAccessToken(existingUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Failed to generate access token",
		})
	}
	// generate refresh token
	refreshToken, err := generateRefreshToken(existingUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Failed to generate refresh token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: !config.AppCfg().Debug,
		Secure:   !config.AppCfg().Debug,
		Expires:  jwt.TimeFunc().Add(time.Duration(config.AppCfg().JWTRefreshKeyExpireHourCount) + time.Hour),
	})

	// Save the refresh token in the database
	tokenRepo := repository.NewTokenRepo(database.GetDB())
	refreshTokenModel := &model.RefreshToken{
		UserID:    existingUser.ID,
		Token:     refreshToken,
		UserAgent: string(c.Request().Header.Peek("User-Agent")),
		IPAddress: c.IP(),
		ExpiresAt: jwt.TimeFunc().Add(time.Duration(config.AppCfg().JWTRefreshKeyExpireHourCount) * time.Hour),
	}

	exists, err := tokenRepo.Exists(existingUser.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Failed to check if refresh token exists " + err.Error(),
		})
	}

	if exists {
		err := tokenRepo.Update(refreshTokenModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": "Failed to update refresh token " + err.Error(),
			})
		}
	} else {
		if err := tokenRepo.Create(refreshTokenModel); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": "Failed to save refresh token ",
			})
		}
	}

	return c.JSON(fiber.Map{
		"msg":          "Login successful",
		"access_token": accessToken,
	})
}

func Logout(c *fiber.Ctx) error {
	tokenRepo := repository.NewTokenRepo(database.GetDB())
	refreshToken := c.Cookies("refresh_token")
	if refreshToken != "" {
		err := tokenRepo.DeleteByToken(refreshToken)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": "Failed to delete refresh token",
			})
		}
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HTTPOnly: !config.AppCfg().Debug,
		Secure:   !config.AppCfg().Debug,
		Expires:  jwt.TimeFunc().Add(-time.Hour), // Set to a past time to invalidate the cookie
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Logout successful",
	})
}

func RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg": "Missing refresh token",
		})
	}

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(config.AppCfg().JWTRefreshKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg": "Invalid refresh token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !claims.VerifyExpiresAt(jwt.TimeFunc().Unix(), true) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg": "Refresh token expired",
		})
	}

	userID := claims["sub"].(string)
	userRepo := repository.NewUserRepo(database.GetDB())
	user, err := userRepo.FindByID(userID)
	if err != nil || user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg": "User not found",
		})
	}

	tokenRepo := repository.NewTokenRepo(database.GetDB())
	refreshTokenModel, err := tokenRepo.FindByToken(refreshToken)
	if err != nil || refreshTokenModel == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg": "Refresh token not found or invalid",
		})
	}

	newAccessToken, err := generateAccessToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Failed to generate new access token",
		})
	}
	// Generate a new refresh token
	newRefreshToken, err := generateRefreshToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Failed to generate new refresh token",
		})
	}

	err = tokenRepo.Update(&model.RefreshToken{
		ID:         refreshTokenModel.ID,
		UserID:     user.ID,
		Token:      newRefreshToken,
		DeviceName: refreshTokenModel.DeviceName,
		IPAddress:  refreshTokenModel.IPAddress,
		ExpiresAt:  jwt.TimeFunc().Add(time.Duration(config.AppCfg().JWTRefreshKeyExpireHourCount) * time.Hour),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Failed to update refresh token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    newRefreshToken,
		HTTPOnly: !config.AppCfg().Debug,
		Secure:   !config.AppCfg().Debug,
		Expires:  jwt.TimeFunc().Add(time.Duration(config.AppCfg().JWTRefreshKeyExpireHourCount) * time.Hour),
	})

	return c.JSON(fiber.Map{
		"access_token": newAccessToken,
	})
}

func generateAccessToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   jwt.TimeFunc().Add(time.Duration(config.AppCfg().JWTAccessKeyExpireMinutesCount) * time.Minute).Unix(),
		"iat":   jwt.TimeFunc().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(config.AppCfg().JWTAccessKey)) // Use a secure key in production
}

func generateRefreshToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  jwt.TimeFunc().Add(time.Duration(config.AppCfg().JWTRefreshKeyExpireHourCount) * time.Hour).Unix(),
		"iat":  jwt.TimeFunc().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(config.AppCfg().JWTRefreshKey)) // Use a secure key in production
}
