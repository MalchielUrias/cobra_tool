# Cobra Tool

Cobra Tool is a simple CLI application built using the [Cobra](https://github.com/spf13/cobra) library. This tool provides basic commands to showcase its functionality, including fetching the app's version and displaying general information about the app.

---

## Features

- **Version Command**: Displays the current version of the application.
- **Info Command**: Provides details about the purpose of the app and its origin.

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/cobra_tool.git
   cd cobra_tool
2. Build the application:

   ```bash
   go build -o cobra_tool

---

## Usage

After building the application, you can use the following commands:

### Version

- Get the current version of the application:

   ```bash
   ./cobra_tool version
   
#### Output 

-  ```lua
   Cobra Tool v1.0 -- HEAD


### Info

- Learn about the app and its purpose:

   ```bash
   ./cobra_tool info
   
#### Output 

-  ```vbnet
   Cobra is a tool that was built by the Hugo team at Golang to help you create CLI applications


---

## Development

1. Install Dependencies:

   ```bash
   go mod tidy

2. Add new commands by creating files and registering them in main.go.


---

## Contributions

Feel free to contribute by submitting issues or pull requests to improve the app.

---

## License
This project is licensed under the MIT License.

---

## Acknowledgments
Cobra Library
Hugo Team