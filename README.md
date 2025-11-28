# Chirpy — HTTPS API Server in Go

**Chirpy** is a REST API server written in Go for managing short messages (*chirps*), user accounts, and subscription status updates via Polka webhooks.

The server supports authentication, user management, metrics, administrative operations, and secure CRUD functionality for chirps.

---

## 🚀 Key Features

* User registration and authentication
* JWT tokens (access + refresh)
* Creating, reading, and deleting *chirps*
* Updating user profiles
* Administrative endpoints
* Polka webhook support (for upgrading user roles)
* Static frontend served via `/app/`

---

# 📡 API Endpoints

Below is a list of **all endpoints** actually present in the project, based on the routing defined in `main.go`.

---

## 🩺 Health & Metrics

### **GET /api/healthz**

Server readiness check.

**Response:**

```json
{ "status": "ok" }
```

---

### **GET /admin/metrics**

Displays the number of file-server hits.

---

### **POST /admin/reset**

Resets server metrics.

---

# 👤 Users

### **POST /api/users**

Create a new user.

**Body:**

```json
{
  "email": "user@example.com",
  "password": "secret"
}
```

---

### **PUT /api/users**

Update the current user's information.
Requires an access token.

**Header:**

```
Authorization: Bearer <token>
```

**Body (example):**

```json
{
  "email": "new@example.com"
}
```

---

### **POST /api/login**

Authenticate a user and issue tokens.

**Response:**

```json
{
  "access_token": "jwt...",
  "refresh_token": "uuid..."
}
```

---

### **POST /api/refresh**

Obtain a new access token using a refresh token.

**Body:**

```json
{ "refresh_token": "uuid..." }
```

---

### **POST /api/revoke**

Revoke a refresh token.

---

# 🐦 Chirps

### **POST /api/chirps**

Create a new chirp.
Requires authorization.

**Body:**

```json
{ "content": "Hello, Chirpy!" }
```

---

### **GET /api/chirps/**

Retrieve a list of all chirps.

**Possible query parameters (if implemented):**

* `author_id`
* `sort=asc|desc`

---

### **GET /api/chirps/{chirpID}**

Retrieve a single chirp by ID.

---

### **DELETE /api/chirps/{chirpID}**

Delete a chirp.
Requires authorization and ownership of the chirp.

---

# 💳 Webhooks

### **POST /api/polka/webhooks**

Webhook endpoint for the Polka service.

Used to update a user’s status (e.g., upgrading to premium).
Triggered by an external system.

---

# 📁 Static Files

### **/app/**

Frontend assets served directly by the file server.

Examples:

```
/app/index.html
/app/assets/logo.png
```

All file-server requests are counted in metrics.

---

# 🔒 Authentication

The server uses:

* **JWT Access Token** — short-lived
* **Refresh Token** — stored in the database, revocable
* Token secret configured via `.env` variable `SECRET`

---

# ⚙️ Environment Variables

The following variables must be set before running the server:

```
DB_URL=<PostgreSQL connection string>
PLATFORM=<environment name>
SECRET=<jwt secret>
```

---

# 📄 License

MIT — free to use and modify.