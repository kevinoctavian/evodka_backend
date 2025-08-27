# Evodka Backend REST API

Proyek ini adalah backend REST API untuk aplikasi Evodka.

## Teknologi yang Digunakan

- **Go** gofiber
- **Database** PostgreSQL
- **Autentikasi JWT**
- **Testing** go test

## Daftar Route dan Kegunaannya

| Method | Endpoint                | Kegunaan                                            |
| ------ | ----------------------- | --------------------------------------------------- |
| POST   | `/api/v1/auth/register` | Registrasi user baru                                |
| POST   | `/api/v1/auth/login`    | Login dan mendapatkan token akses dan token refresh |
| POST   | `/api/v1/auth/refresh`  | mendapatkan kembali token akses                     |
| POST   | `/api/v1/auth/logout`   | menghapus semua akses user                          |

> **Catatan:** Daftar route di atas dapat disesuaikan dengan kebutuhan aplikasi.

## Cara Kerja Setiap Route

- **/auth/register**: Menerima data user baru dan menyimpannya ke database.
- **/auth/login**: Memvalidasi kredensial dan mengembalikan JWT token.
- **/users/me**: Mengambil data user yang sedang login menggunakan token.
- **/items/**: Mengelola data item (CRUD: Create, Read, Update, Delete).

---
