<h1 align="center">GoWebBase</h1>
<p align="center">
  <img src="https://socialify.git.ci/MostafaSensei106/GoWebBase/image?custom_language=Go&font=KoHo&language=1&logo=https%3A%2F%2Favatars.githubusercontent.com%2Fu%2F138288138%3Fv%3D4&name=1&owner=1&pattern=Floating+Cogs&theme=Light" alt="Go App Banner">
  <br>
  <strong>A simple Go web service with a modern user interface.</strong><br>
  This project demonstrates a basic yet functional web application built with Go, serving static files, handling form submissions, and providing a foundation for further development.
</p>

<p align="center">
  <a href="#-description">Description</a> •
  <a href="#-features">Features</a> •
  <a href="#-architecture">Architecture</a> •
  <a href="#-technologies">Technologies</a> •
  <a href="#-getting-started">Getting Started</a> •
  <a href="#-screenshots">Screenshots</a> •
  <a href="#-api-endpoints">API Endpoints</a> •
  <a href="#-contributing">Contributing</a> •
  <a href="#-license">License</a>
</p>

---

## 📝 Description

GoWebBase is a lightweight Go application designed to provide a foundational web service with a simple yet modern user interface. It focuses on demonstrating core web development concepts in Go, including HTTP server setup, static asset serving, and basic form data processing. The UI has been enhanced with modern CSS and client-side JavaScript for input validation, offering a better user experience.

---

## ✨ Features

### Core Functionality

- **Static File Serving**: Efficiently serves HTML, CSS, and JavaScript files from the `/static` directory.
- **Basic Routing**: Implements distinct routes for different functionalities (`/hello`, `/form`).
- **Form Handling**: Processes `POST` requests from an HTML form, demonstrating data capture from user input.
- **Modern User Interface**: Styled with custom CSS for a clean and responsive design, supporting a better aesthetic.
- **Client-Side Validation**: Includes JavaScript for basic input validation on the form page, enhancing user experience and providing immediate feedback.

---

## 🏗️ Architecture

The application is structured to maintain a clear separation of concerns, making it easy to understand, extend, and maintain.

### Project Structure

```
.
├── constants/        # Global constants and route definitions
├── data/             # Data models (e.g., user_model.go)
├── logic/            # Business logic and request handlers
├── main.go           # Application entry point
├── server/           # HTTP server setup and route registration
└── static/           # Static web assets (HTML, CSS, JS)
    ├── css/          # Custom stylesheets
    ├── js/           # Client-side JavaScript for interactions and validation
    ├── form.html     # HTML page for form submission
    └── index.html    # Main landing HTML page
```

- **`main.go`**: Initializes and starts the HTTP server.
- **`server/`**: Contains the `Execute` function which sets up all the application's routes and starts the `http.ListenAndServe`.
- **`constants/`**: Defines global constants, including route paths to ensure consistency across the application.
- **`logic/handlers/`**: Houses the Go functions responsible for handling specific HTTP requests (e.g., `FormHandler`, `HelloHandler`).
- **`static/`**: This directory holds all the web-facing assets like HTML templates, custom CSS, and client-side JavaScript files.

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

## 🚀 Getting Started

To get the GoWebBase application up and running on your local machine, follow these steps.

### Prerequisites

- **Go**: Version `1.25.5` or higher.
  - Verify your Go installation: `go version`

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/MostafaSensei106/GoWebBase.git
    cd GoWebBase
    ```

2.  **Tidy Go modules:**
    This command ensures that your `go.mod` and `go.sum` files are up to date and that all necessary dependencies (if any external were introduced) are downloaded.
    ```bash
    go mod tidy
    ```

### Running the Application

1.  **Execute the main program:**

    ```bash
    go run main.go
    ```

    The server will start and typically listen on `http://localhost:8080`. You will see output like `Starting server at port 8080`.

2.  **Access in Browser:**
    - **Home Page**: Open your web browser and go to `http://localhost:8080/static/index.html`
    - **Form Page**: Navigate to `http://localhost:8080/static/form.html`
    - **Hello Endpoint**: Visit `http://localhost:8080/hello`

---

## 🖼️ Screenshots

| Home Page                                                                                 | Form Page                                                                               |
| :---------------------------------------------------------------------------------------- | :-------------------------------------------------------------------------------------- |
| ![Home Page](screenshots/Screenshot%202025-12-04%20at%2001-52-38%20GoWebBase%20-%20Home.png) | ![Form Page](screenshots/Screenshot%202025-12-04%20at%2001-52-48%20GoWebBase%20-%20Form.png) |

---

## 🔗 API Endpoints

The GoWebBase application exposes the following HTTP endpoints:

- **`GET /static/{file}`**

  - **Description**: Serves static files (e.g., `index.html`, `form.html`, `style.css`, `form.js`) directly from the `static` directory.
  - **Example**: `GET /static/index.html`

- **`GET /hello`**

  - **Description**: A simple endpoint that responds with a "Hello!" message.
  - **Response**: `hello` (plain text)

- **`GET /form`**

  - **Description**: Renders and serves the `form.html` page, which contains the user input form.

- **`POST /form`**
  - **Description**: Processes the data submitted via the HTML form. It expects `name` and `address` parameters from the form.
  - **Parameters**:
    - `name` (string): The name provided by the user.
    - `address` (string): The address provided by the user.

---

## 🤝 Contributing

Contributions to the GoWebBase project are welcome! If you have suggestions for improvements, new features, or bug fixes, please feel free to:

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
