<h1 align="center">GoWebBase</h1>
<p align="center">
  <img src="https://socialify.git.ci/MostafaSensei106/GoWebBase/image?custom_language=Go&font=KoHo&language=1&logo=https%3A%2F%2Favatars.githubusercontent.com%2Fu%2F138288138%3Fv%3D4&name=1&owner=1&pattern=Floating+Cogs&theme=Light" alt="GoWebBase Banner">
</p>

<p align="center">
  <strong>A simple Go web service with a modern user interface.</strong><br>
  A foundational project demonstrating a web application in Go, serving static files and handling form submissions.
</p>

<p align="center">
  <a href="#-about">About</a> •
  <a href="#-features">Features</a> •
  <a href="#-installation">Installation</a> •
  <a href="#-technologies">Technologies</a> •
  <a href="#-contributing">Contributing</a> •
  <a href="#-license">License</a>
</p>

---

## 📝 About

Welcome to **GoWebBase** — a lightweight and easy-to-understand web application starter kit built with Go. GoWebBase empowers developers with a clean and logical project structure for building web services. It provides a foundational example of serving static content, handling API routes, and processing user input, all with a modern frontend built with simple HTML, CSS, and JavaScript.

---

## ✨ Features

### Core Functionality

- **Static File Serving**: Efficiently serves HTML, CSS, and JavaScript files from the `/static` directory.
- **Basic Routing**: Implements distinct routes for different functionalities (`/hello`, `/form`).
- **Form Handling**: Processes `POST` requests from an HTML form, demonstrating data capture from user input.
- **Modern User Interface**: Styled with custom CSS for a clean and responsive design, supporting a better aesthetic.
- **Client-Side Validation**: Includes JavaScript for basic input validation on the form page, enhancing user experience and providing immediate feedback.
- **Cross-Platform Builds**: The included `Makefile` supports building executables for Linux and Windows across multiple architectures (386, amd64, arm, arm64).

---

## 📦 Installation

You can get the GoWebBase application up and running in two ways: by downloading a pre-built release or by building it from the source code.

### Easy Install (from Releases)

Download the latest pre-built and packaged binary for your platform from the [Releases](https://github.com/MostafaSensei106/GoWebBase/releases) page.

#### 🐧 Linux

1.  Download the appropriate `gowebbase-vX.Y.Z-linux-[arch].tar.gz` archive.
2.  Extract the archive:
    ```bash
    tar -xzf gowebbase-vX.Y.Z-linux-[arch].tar.gz
    ```
3.  Navigate into the extracted directory:
    ```bash
    cd linux/[arch]
    ```
4.  Run the application:
    ```bash
    ./gowebbase
    ```

---

#### 🪟 Windows

1.  Download the appropriate `gowebbase-vX.Y.Z-windows-[arch].zip` archive.
2.  Extract the archive using Windows Explorer or your favorite tool.
3.  Open PowerShell or Command Prompt, navigate into the extracted directory:
    ```powershell
    cd windows/[arch]
    ```
4.  Run the application:
    ```powershell
    .\gowebbase.exe
    ```

---

### 🏗️ Build from Source

> ![Note]
> GoWebBase uses a `Makefile` to simplify the build process.
> Make sure you have `make`, `Go`, and `git` installed on your system.

#### Step 1: Install Build Tools (if not already installed)

##### For **Arch Linux** and based distros:

```bash
sudo pacman -S base-devel mingw-w64-gcc
```

##### For **Debian / Ubuntu** and based distros:

```bash
sudo apt install build-essential gcc-mingw-w64-x86-64
```

##### For **Fedora** and based distros:

```bash
sudo dnf install make mingw64-gcc
```

##### For **Windows**:

- Install [MSYS2](https://www.msys2.org/) (recommended) or use [Git Bash](https://gitforwindows.org/).
- Inside the shell, install make:
  ```bash
  pacman -Syu
  pacman -S make
  ```

---

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
    Run the `make build` command. This will automatically check and install dependencies, format the code, and build the executable for your current OS.
    ```bash
    make build
    # Or, as an alias:
    make all
    ```
    To create release archives for all supported platforms (Linux/Windows), run:
    ```bash
    make release
    ```

---

#### ✅ Result

- `make build` will compile `GoWebBase` and place the executable and the `static` folder in `bin/<your-os>/<your-arch>`.
- `make release` will create distributable `.zip` and `.tar.gz` archives in the `release/` directory.

---

## 🚀 Running the Application

After building or extracting the application, you can run it.

1.  **Execute the main program:**
    Navigate to the directory containing the executable and the `static` folder, and run it.

    **On Linux:**

    ```bash
    ./gowebbase
    ```

    **On Windows:**

    ```powershell
    .\gowebbase.exe
    ```

    The server will start and listen on `http://localhost:8080`.

2.  **Access in Browser:**
    - **Home Page**: `http://localhost:8080/static/index.html`
    - **Form Page**: `http://localhost:8080/static/form.html`
    - **Hello Endpoint**: `http://localhost:8080/hello`

---

## 🏗️ Architecture

The application is structured to maintain a clear separation of concerns, making it easy to understand, extend, and maintain.

### Project Structure

```
.
├── Makefile          # Automates build, release, and clean tasks
├── bin/              # Output for compiled binaries
├── constants/        # Global constants and route definitions
├── data/             # Data models
├── logic/            # Business logic and request handlers
├── main.go           # Application entry point
├── release/          # Output for release archives
├── server/           # HTTP server setup and route registration
└── static/           # Static web assets (HTML, CSS, JS)
```

---

## 🛠️ Technologies

This project leverages the following technologies:

| Technology        | Description                                                         |
| :---------------- | :------------------------------------------------------------------ |
| 🐦 **Go**         | The core programming language for the backend web service.          |
| 🖥️ **HTML5**      | Standard markup language for structuring web content.               |
| 🎨 **CSS3**       | Styling language used for the modern and responsive user interface. |
| 📜 **JavaScript** | Client-side scripting for interactive elements and form validation. |

---

## 🖼️ Screenshots

| Home Page                                                                                    | Form Page                                                                                    |
| :------------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------------------- |
| ![Home Page](screenshots/Screenshot%202025-12-04%20at%2001-52-38%20GoWebBase%20-%20Home.png) | ![Form Page](screenshots/Screenshot%202025-12-04%20at%2001-52-48%20GoWebBase%20-%20Form.png) |

---

## 🔗 API Endpoints

The GoWebBase application exposes the following HTTP endpoints:

- **`GET /static/{file}`**
  - **Description**: Serves static files from the `static` directory.
- **`GET /hello`**
  - **Description**: A simple endpoint that responds with a "Hello!" message.
- **`POST /form`**
  - **Description**: Processes the data submitted via the HTML form.

---

## 🤝 Contributing

Contributions to the GoWebBase project are welcome!

1.  Fork the Repository
2.  Create your Feature Branch (`git checkout -b feature/your-feature-name`)
3.  Commit your Changes (`git commit -m 'feat: Add some amazing feature'`)
4.  Push to the Branch (`git push origin feature/your-feature-name`)
5.  Open a Pull Request

---

## 📜 License

This project is open-source and available under the **MIT License**.

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/MostafaSensei106">MostafaSensei106</a>
</p>
