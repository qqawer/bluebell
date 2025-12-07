## 🚀 Project Title

**Bluebell** is a community-based web application built with Go and the Gin framework. It allows users to sign up, log in, create posts within communities, and vote on content. It aims to provide a high-performance backend for a Reddit-like platform.

## ✨ Features

### User Experience

  * **User Authentication**: Secure sign-up and login functionality using JWT.
  * **Community Interaction**: Browse different communities and view their details.
  * **Post Management**: Create new posts, view post details, and list posts by time or score.
  * **Voting System**: Vote on posts to influence their ranking.

### APIs

  * **Public Endpoints**: User registration and login.
  * **Protected Endpoints**: All community, post, and voting operations require JWT authentication.
  * **Swagger Documentation**: Integrated Swagger UI for API exploration.

-----

## 📐 Architecture and Technologies

  * **Backend Framework:** **Gin** (Go Web Framework).
  * **Language & Runtime:** **Go 1.22.9**.
  * **Persistence:**
    * **MySQL**: Primary relational database for user and content data (using **GORM**).
    * **Redis**: Used for caching and potentially ranking logic.
  * **Authentication:** **JWT (JSON Web Tokens)** for stateless authentication.
  * **Configuration:** **Viper** for managing configuration files.
  * **Logging:** **Zap** for structured, high-performance logging.
  * **ID Generation:** **Snowflake** algorithm for generating unique distributed IDs.

-----

## 📁 Project Layout

```text
.
├── conf/                 # Configuration files (config.yaml)
├── controllers/          # API handlers (MVC controllers)
├── dao/                  # Data Access Objects (MySQL, Redis)
├── docs/                 # Swagger documentation files
├── logger/               # Logging setup (Zap)
├── logic/                # Business logic layer
├── middlewares/          # Gin middlewares (Auth, CORS, etc.)
├── models/               # Data models (Structs)
├── pkg/                  # Shared packages (Snowflake, JWT, etc.)
├── router/               # Route definitions
├── settings/             # Configuration loading logic
├── main.go               # Application entry point
├── Makefile              # Build and run commands
└── go.mod                # Go module dependencies
```

-----

## 🔧 Prerequisites

  * **Go**: Version 1.22.9 or newer.
  * **MySQL**: Running on `localhost:3306` (or configured host).
  * **Redis**: Running on `localhost:6379` (or configured host).

-----

## ⚙️ Configuration

The primary configuration is in `conf/config.yaml`.

### Example Configuration

```yaml
app:
  name: "web_app"
  port: 8080
  mode: "release"

mysql:
  host: "mysql"
  port: 3306
  user: "root"
  password: "123456"
  name: "webApp"

redis:
  host: "redis"
  port: 6379
```

> **Security Note:** Avoid committing sensitive passwords to version control. Use environment variables or a separate secret management solution in production.

-----

## 🛠️ Setup and Running

### Database Bootstrap

1.  **Create the Database**:
    Ensure a MySQL database named `webApp` exists (or update `conf/config.yaml` to match your DB name).
    ```sql
    CREATE DATABASE IF NOT EXISTS webApp;
    ```
2.  **Initialize Tables**:
    The application uses GORM, which may auto-migrate schemas, or you can use the provided `init.sql` if available (checked file list, `init.sql` exists).
    ```bash
    mysql -u root -p webApp < init.sql
    ```

### Running the Application

1.  **Build the project**:
    ```bash
    make build
    ```
    This will create a binary named `bluebell` in the `bin/` directory.

2.  **Run the application**:
    ```bash
    make run
    # Or run the binary directly:
    ./bin/bluebell conf/config.yaml
    ```

The application will start on port `8080` (default).

-----

## 🧪 Testing

Run the Go tests using the standard toolchain:

```bash
go test ./...
```

Or use the Makefile target for formatting and vetting:

```bash
make gotool
```

-----

## 🔗 HTTP Entry Points / API Reference

| Audience | Method | Path | Purpose |
|----------|--------|------|---------|
| Public   | POST   | `/api/v1/signup` | User registration. |
| Public   | GET    | `/api/v1/login` | User login (returns JWT). |
| Auth     | GET    | `/api/v1/community` | List all communities. |
| Auth     | GET    | `/api/v1/community/:id` | Get community details. |
| Auth     | POST   | `/api/v1/post` | Create a new post. |
| Auth     | GET    | `/api/v1/post/:id` | Get post details. |
| Auth     | GET    | `/api/v1/posts` | Get post list. |
| Auth     | POST   | `/api/v1/vote` | Vote on a post. |

> **Note:** Swagger documentation is available at `/swagger/index.html` when the application is running.

-----

## 🛑 Known Limitations and Future Ideas

  * **Login Method:** The login endpoint uses `GET`, which is unconventional. Consider changing to `POST` for better security practice.
  * **Configuration:** Database passwords are currently in `config.yaml`. Move to environment variables.
  * **Testing:** More comprehensive unit and integration tests can be added.

-----

## 🤝 Contributors

  * **Bluebell Team**
