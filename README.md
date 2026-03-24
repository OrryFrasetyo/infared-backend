# 🚑 InfaRed API Documentation (MVP)

InfaRed (Sistem Logistik Bencana Cerdas) Backend API. Dibangun menggunakan **Golang (Gin + SQLX)**, **PostgreSQL**, dan terintegrasi dengan **Google Gemini AI**.

---

## 🌍 Base URL

- **Local:** `http://localhost:8080/api/v1`
- **Production:** `(Akan diisi URL Cloud jika sudah deploy)`

## 🛡️ Authentication

Sebagian besar _endpoint_ membutuhkan JWT Token. Tambahkan token pada _Header_ saat melakukan _request_:
`Authorization: Bearer <token_jwt_anda>`

---

## 📦 Standard Response Format

Sistem ini menggunakan standar _wrapper_ JSON untuk semua _response_.

**✅ Success Response (HTTP 200 / 201)**

```json
{
    "success": true,
    "message": "Deskripsi aksi berhasil",
    "data": { ... } // Bisa berupa object atau array
}
```

**❌ Error Response (HTTP 400 / 401 / 403 / 500)**

```json
{
  "success": false,
  "message": "Pesan error utama (human-readable)",
  "errors": "Detail teknis error (jika ada)"
}
```

## 📌 1. Authentication & Users

### Login

- **Endpoint:** POST /auth/login
- **Auth:** ❌ Public
- **Request Body:**
  ```json
  {
    "email": "admin@infared.com",
    "password": "password123"
  }
  ```
- **Response Data:** Mengembalikan token dan object user (id, name, email, role).

### Register Admin/Relawan

- **Endpoint:** POST /auth/register
- **Auth:** 🔒 Admin Only
- **Request Body:** name, email, password (min 6 char).

### Get All Users

- **Endpoint:** GET /users
- **Auth:** 🔒 Admin Only
- **Response Data:** Array of user objects.

## 📌 2. Master Data (Posko & Barang)

### Get All Posko

- **Endpoint:** GET /posko
- **Auth:** 🔒 Token Required (Admin & Relawan)
- **Response Data:** Array of posko objects.

### Get Posko Inventory

- **Endpoint:** GET /posko/:id/inventory
- **Auth:** 🔒 Token Required
- **Response Data:** Array of inventory objects dilengkapi dengan item_name dan item_unit.

### Get All Items (Barang)

- **Endpoint:** GET /items
- **Auth:** 🔒 Token Required
- **Response Data:** Array of item objects.

### Create Item

- **Endpoint:** POST /items
- **Auth:** 🔒 Admin Only
- **Request Body:**
  ```json
  {
    "name": "Selimut Tebal",
    "unit": "Pcs"
  }
  ```

## 📌 3. AI Logistics Core (Transaksi)

### Submit Chat to AI (Lapor Logistik)

- **Endpoint:** POST /requests/chat
- **Auth:** 🔒 Token Required
- **Description:** Mengirim teks natural relawan ke Gemini AI untuk diekstrak menjadi struktur data barang dan tingkat urgensinya.
- **Request Body:**
  ```json
  {
    "posko_id": "psk-12345",
    "prompt_text": "Tolong segera kirimkan 50 selimut tebal dan 10 dus air mineral! Badai merusak tenda."
  }
  ```
- **Response Data:** Mengembalikan object request beserta array items hasil ekstraksi AI yang sudah disimpan ke database.

### Get All Logistics Requests (Dashboard)
- **Endpoint:** GET /requests
- **Auth:** 🔒 Token Required
- **Description:** Menampilkan daftar riwayat permintaan logistik beserta detail barang, nama posko, dan nama pelapor (hasil JOIN SQL).
- **Response Data:** Array of request objects dengan relasi bersarang (nested arrays).
