# backend-fleetifyid-challenge

Ini adalah backend untuk Fullstack Developer Challenge Test Fleetify.id, disini saya menggunakan golang sebagai backend dan menggunakan Go Fiber, godotenv, serta Gorm.

## Folder Structure

Ini adalah struktur folder API.

```bash
backend-fleetifyid-challenge/
│
├─ config/
│   └─ database.go          # Koneksi database
│
├─ controllers/
│   ├─ attendancecontroller/
│   │   └─ attendance.go    # Controller Absensi
│   ├─ departementcontroller/
│   │   └─ departement.go   # Controller Departemen
│   └─ employeecontroller/
│       └─ employee.go      # Controller Karyawan
│
├─ models/
│   ├─ attendance_history.go # Model attendance_history
│   ├─ attendance.go         # Model attendance
│   ├─ departement.go        # Model departement
│   └─ employee.go           # Model employee
│
├─ .env                      # Konfigurasi environment
├─ example.env               # Contoh file environment
├─ db.sql                    # File SQL untuk setup database
├─ go.mod
├─ go.sum
├─ main.go                   # File utama api
├─ LICENSE
└─ README.md

```

## Installation

### 1. Clone Repository

```bash
git clone https://github.com/DhikaNino/backend-fleetifyid-challenge.git
cd backend-fleetifyid-challenge
```

### 2. Install Dependency

```bash
go mod tidy
```

### 3. Setup Database

Import file db.sql yang sudah ada di projek:

```bash
mysql -u root -p db_fleetifyid_challenge < db.sql
```

### 4. Setup Environment

Rename exmaple.env dan ubah konfigurasinya dan sesuakan atau buat file .env di root folder projek dan isi dengan konfigurasi berikut:

```bash
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=db_fleetifyid_challenge

API_PORT=3000
```

## Run API

Jalankan ini untuk memulai API:

```python
go run main.go
```

## API Request

### Base URL:

```bash
http://localhost:3000/api
```

### 1. Employee API

#### 1.1 Get All Employees

```bash
Endpoint: GET /employee
```

Mendapatkan semua karyawan.

Response:

```json
{
  "data": [
    {
      "id": 1,
      "employee_id": "KRY-00001",
      "departement_id": 1,
      "name": "Andhika",
      "address": "Bogor",
      "created_at": "2025-08-16T00:24:50Z",
      "updated_at": "2025-08-16T00:24:50Z",
      "departement": {
        "id": 1,
        "departement_name": "Human Resource Departement",
        "max_clock_in_time": "08:00:00",
        "max_clock_out_time": "17:00:00"
      }
    }
  ],
  "meta": {
    "last_page": 1,
    "limit": 10,
    "page": 1,
    "search": "",
    "total": 1
  }
}
```

#### 1.2 Get Single Employee

```bash
Endpoint: GET /employee/:employee_id
```

Mendapatkan data karyawan berdasarkan employee_id.

Response:

```json
{
  "data": {
    "id": 1,
    "employee_id": "KRY-00068",
    "departement_id": 1,
    "name": "Andhika",
    "address": "Bogor",
    "created_at": "2025-08-16T00:26:29Z",
    "updated_at": "2025-08-16T00:26:29Z",
    "departement": {
      "id": 1,
      "departement_name": "Human Resource Departement",
      "max_clock_in_time": "08:00:00",
      "max_clock_out_time": "17:00:00"
    }
  }
}
```

#### 1.3 Create Employee

```bash
Endpoint: POST /employee
```

Menambahkan karyawan baru.

Body (JSON):

```json
{
  "departement_id": 1,
  "name": "Andhika",
  "address": "Bogor"
}
```

#### 1.4 Update Employee

```bash
Endpoint: PUT /employee/:employee_id
```

Mengubah data karyawan berdasarkan employee_id.

Body (JSON):

```json
{
  "departement_id": 1,
  "name": "Andhika",
  "address": "Bogor"
}
```

#### 1.5 Delete Employee

```bash
Endpoint: DELETE /employee/:id
```

Menghapus data employee berdasarkan id.

### 2. Employee API

#### 2.1 Get All Departement

```bash
Endpoint: GET /departement
```

Mendapatkan semua departement.

Response:

```json
{
  "data": [
    {
      "id": 1,
      "departement_name": "Human Resource Departement",
      "max_clock_in_time": "08:00:00",
      "max_clock_out_time": "17:00:00"
    }
  ],
  "meta": {
    "last_page": 1,
    "limit": 10,
    "page": 1,
    "search": "",
    "total": 1
  }
}
```

#### 2.2 Get Single Departement

```bash
Endpoint: GET /departement/:id
```

Mendapatkan data departement berdasarkan id.

Response:

```json
{
  "id": 1,
  "departement_name": "Human Resource Departement",
  "max_clock_in_time": "08:00:00",
  "max_clock_out_time": "17:00:00"
}
```

#### 2.3 Create Departement

```bash
Endpoint: POST /departement
```

Menambahkan departement baru.

Body (JSON):

```json
{
  "departement_name": "Human Resource Departement",
  "max_clock_in_time": "08:00:00",
  "max_clock_out_time": "17:00:00"
}
```

#### 2.4 Update Departement

```bash
Endpoint: PUT /departement/:id
```

Mengubah data departement berdasarkan id.

Body (JSON):

```json
{
  "departement_name": "Human Resource Departement",
  "max_clock_in_time": "08:00:00",
  "max_clock_out_time": "17:00:00"
}
```

#### 2.5 Delete Departement

```bash
Endpoint: DELETE /departement/:id
```

Menghapus data departement berdasarkan id.

### 3. Atendance API

#### 3.1 Get Attendance Logs

Mendapatkan data attendance logs.

```bash
Endpoint: GET /attendance
```

Mendapatkan data attendance logs dengan filter berdasarkan tanggal dan departemen.

```bash
Endpoint: GET /attendance?departement_id=11&start_date=2025-08-16&end_date=2026-01-03
```

Parameter:

```bash
departement_id: 1
start_date: 2025-08-16
end_date: 2026-01-03
```

Response:

```json
{
  "data": [
    {
      "employee_id": "KRY-00001",
      "employee_name": "Andhika",
      "departement_name": "Human Resource Departement",
      "date_attendance": "2025-08-16T03:14:55Z",
      "attendance_type": 1,
      "description": "Absen masuk",
      "max_clock_in_time": "08:00:00",
      "max_clock_out_time": "17:00:00",
      "status_ketepatan": "Terlambat"
    },
    {
      "employee_id": "KRY-00001",
      "employee_name": "Andhika",
      "departement_name": "Human Resource Departement",
      "date_attendance": "2025-08-16T03:42:32Z",
      "attendance_type": 2,
      "description": "Absen pulang",
      "max_clock_in_time": "08:00:00",
      "max_clock_out_time": "17:00:00",
      "status_ketepatan": "Pulang lebih awal"
    }
  ]
}
```

#### 3.2 Atendance In

```bash
Endpoint: POST /attendance/in
```

Absensi masuk karyawan

Body (JSON):

```json
{
  "employee_id": "KRY-00001"
}
```

#### 3.3 Atendance Out

```bash
Endpoint: PUT /attendance/out
```

Absensi keluar karyawan

Body (JSON):

```json
{
  "employee_id": "KRY-00001"
}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
