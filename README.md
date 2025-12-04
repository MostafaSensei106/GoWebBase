<h1 align="center">GoWebBase</h1>
<p align="center">
  <img src="https://socialify.git.ci/MostafaSensei106/GoWebBase/image?custom_language=Go&font=KoHo&language=1&logo=https%3A%2F%2Favatars.githubusercontent.com%2Fu%2F138288138%3Fv%3D4&name=1&owner=1&pattern=Floating+Cogs&theme=Light" alt="GoWebBase Banner">
</p>

<p align="center">
  <strong>A simple Go web service with a modern user interface.</strong><br>
  A foundational boilerplate for building web applications in Go, demonstrating a clean project structure, static file serving, and dynamic request handling.
</p>

<p align="center">
  <a href="#-about">About</a> •
  <a href="#-features">Features</a> •
  <a href="#-architecture">Architecture</a> •
  <a href="#-installation">Installation</a> •
  <a href="#-running-the-application">Running</a> •
  <a href="#-makefile-commands">Makefile Commands</a> •
  <a href="#-api-endpoints">API Endpoints</a> •
  <a href="#-contributing">Contributing</a>
</p>

---

## 📝 About

Welcome to **GoWebBase** — a lightweight and easy-to-understand web application starter kit built with Go. This project is designed to serve as a robust foundation (or boilerplate) for developers looking to build web services and applications.

It demonstrates core concepts of web development in Go, including a clean and scalable project structure, HTTP server setup, static asset serving, and dynamic API route handling. The frontend is intentionally simple, built with modern HTML, CSS, and vanilla JavaScript to showcase backend functionality without the complexity of a large frontend framework. This makes it an ideal starting point for a wide range of projects, from simple APIs to more complex, full-stack web applications.

---

## ✨ Features

-   **Clean & Scalable Architecture**: The project is organized with a clear separation of concerns, making it easy to extend and maintain as your application grows.
-   **Static File Serving**: Efficiently serves a modern frontend built with HTML, CSS, and JavaScript from the `/static` directory.
-   **Dynamic API Routing**: Includes examples for basic API endpoints (`/hello`) and form handling (`/form`), showcasing how to process different HTTP methods.
-   **Modern UI with Validation**: The frontend includes a responsive and clean user interface with client-side form validation for a better user experience.
-   **Robust Build System**: A comprehensive `Makefile` automates common development tasks like building, checking, and cleaning the project.
-   **Cross-Platform Releases**: The `Makefile` can produce self-contained, cross-platform release archives for Linux and Windows across `386`, `amd64`, `arm`, and `arm64` architectures.
-   **Docker Support**: Comes with a multi-stage `Dockerfile` and `Makefile` targets for easy containerization and deployment.
-   **CI/CD Ready**: Includes a GitHub Actions workflow to automatically build and package the application on every push or pull request to the `main` branch.

---

## 🏗️ Architecture

The application is structured to be modular and maintainable. It follows a logical separation of concerns, where different parts of the application have distinct responsibilities.

### Project Structure

```
.
├── .github/
│   └── workflows/
│       └── build-and-release.yml  # GitHub Actions CI/CD workflow
├── Dockerfile                  # For building a containerized version of the app
├── Makefile                    # Automates build, release, and clean tasks
├── README.md                   # Project documentation
├── bin/                        # Output for compiled binaries (created by make)
├── constants/
│   ├── data.go                 # (Placeholder for data-related constants)
│   └── routs.go                # Defines API route paths as constants
├── data/
│   └── models/
│       └── user_model.go       # Defines the User struct
├── go.mod                      # Go module definition
├── logic/
│   └── handlers/
│       ├── form_handler.go     # Logic for handling POST /form requests
│       └── hello_handler.go    # Logic for handling GET /hello requests
├── main.go                     # Main application entry point
├── release/                    # Output for release archives (created by make release)
├── server/
│   └── routes.go               # HTTP server setup and route registration
└── static/
    ├── css/
    │   └── style.css           # Custom CSS for the frontend
    ├── js/
    │   └── form.js             # Client-side form validation
    ├── form.html               # HTML page for the form
    └── index.html              # Main landing HTML page
```

### Request Flow

A typical HTTP request flows through the application as follows:

1.  **Entry Point**: The application is started by running `go run main.go`. The `main()` function in `main.go` makes a single call to `server.Execute()`.
2.  **Server Setup**: The `Execute()` function in `server/routes.go` is responsible for all setup. It initializes the HTTP server, registers all route handlers, and starts the server to listen for requests on port `8080`.
3.  **Routing**: The `http.Handle` and `http.HandleFunc` calls map URL paths (defined in `constants/routs.go`) to specific handler functions located in `logic/handlers/`.
    -   Requests for static files (e.g., `/static/index.html`) are managed by a `http.FileServer`.
    -   Requests for dynamic routes (e.g., `/hello`, `/form`) are passed to their corresponding handler functions.
4.  **Handling Logic**: The handler function (e.g., `hello_handler.go`) processes the request and writes a response back to the `http.ResponseWriter`. For the form, `form_handler.go` parses the submitted form data from the request body.

---

## 📦 Installation

You can get the GoWebBase application up and running in two ways: by downloading a pre-built release or by building it from the source code.

### Easy Install (from Releases)

Download the latest pre-built and packaged binary for your platform from the [Releases](https://github.com/MostafaSensei106/GoWebBase/releases) page.

#### 🐧 Linux
1.  Download `gowebbase-vX.Y.Z-linux-[arch].tar.gz`.
2.  Extract the archive: `tar -xzf gowebbase-vX.Y.Z-linux-[arch].tar.gz`
3.  Navigate into the directory: `cd linux/[arch]`
4.  Run the application: `./gowebbase`

#### 🪟 Windows
1.  Download `gowebbase-vX.Y.Z-windows-[arch].zip`.
2.  Extract the archive.
3.  Open a terminal and navigate into the directory: `cd windows/[arch]`
4.  Run the application: `.\gowebbase.exe`

---

### 🏗️ Build from Source

> ![Note]
> GoWebBase uses a `Makefile` to simplify the build process. Make sure you have `make`, `Go` (version 1.25.5), and `git` installed. On Windows, a POSIX-compliant shell like Git Bash or MSYS2 is recommended.

#### Step 1: Install Build Tools (Optional, for cross-compilation)
To build for Windows from a non-Windows OS, you may need a cross-compiler.

-   **Arch Linux**: `sudo pacman -S base-devel mingw-w64-gcc`
-   **Debian/Ubuntu**: `sudo apt install build-essential gcc-mingw-w64-x86-64`
-   **Fedora**: `sudo dnf install make mingw64-gcc`

#### Step 2: Clone and Build
1.  **Clone the repository:**
    ```bash
    git clone --depth 1 https://github.com/MostafaSensei106/GoWebBase.git
    ```
2.  **Navigate to the project directory:**
    ```bash
    cd GoWebBase
    ```
3.  **Build the application:**
    This command automatically checks dependencies, formats code, and builds the executable.
    ```bash
    make build
    ```

---

## 🚀 Running the Application

### Natively
After building or extracting the application, navigate to the directory containing the executable and the `static` folder, and run it.

-   **On Linux**: `./gowebbase`
-   **On Windows**: `.\gowebbase.exe`

### 🐳 Running with Docker
If you have Docker installed, you can easily build and run the application in a container.

1.  **Build and run the container:**
    ```bash
    make docker-run
    ```
2.  **Access the application** in your browser at `http://localhost:8080/static/index.html`.

---

## ⚙️ Makefile Commands

The `Makefile` provides several commands to streamline development and builds:

| Command           | Description                                                                        |
| :---------------- | :--------------------------------------------------------------------------------- |
| `make all`        | An alias for `make build`. The default command.                                    |
| `make build`      | Builds the executable and copies static files for the current OS/architecture.     |
| `make release`    | Builds and creates compressed release archives for all target platforms.           |
| `make install`    | A convenience alias for `make build`. Does not install the binary system-wide.     |
| `make docker-build`| Builds the Docker image for the application.                                       |
| `make docker-run` | Builds and runs the application inside a Docker container.                           |
| `make check`      | Runs all code quality checks (`deps`, `fmt`, `vet`).                                 |
| `make deps`       | Checks and verifies Go module dependencies.                                        |
| `make fmt`        | Formats all Go source files in the project.                                        |
| `make vet`        | Runs `go vet` to report suspicious constructs in the code.                           |
| `make clean`      | Deletes all build artifacts, release archives, and Go caches.                      |
| `make help`       | Shows a help message documenting all available commands.                           |

---

## 🔗 API Endpoints

-   **`GET /static/{file}`**
    -   **Description**: Serves static files from the `static` directory.
-   **`GET /hello`**
    -   **Description**: A simple endpoint to check if the server is running.
    -   **Example**: `curl http://localhost:8080/hello`
    -   **Response**: `hello`
-   **`POST /form`**
    -   **Description**: Processes form data submitted from the `/static/form.html` page.
    -   **Example**:
        ```bash
        curl -X POST -F "name=John Doe" -F "address=123 Main St" http://localhost:8080/form
        ```
    -   **Response**: Prints the submitted name and address to the server console.

---

## 🤝 Contributing

Contributions are welcome! Please feel free to fork the repository, make changes, and open a pull request.

---

## 📜 License

This project is open-source and available under the **MIT License**.

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/MostafaSensei106">MostafaSensei106</a>
</p>