GPM - Go Package Manager
GPM is a Go Package Manager designed to simplify and streamline the process of fetching and managing Go packages. Inspired by go get, GPM provides additional customization and potentially more user-friendly features, aimed at improving the Go development experience.

Features
Simple Package Installation: Fetch and install Go packages with a single command.
Package Mapping: Easily manage package mappings through a predefined list, eliminating the need for complex search queries.
Automated Package Downloading: Automatically download and set up Go packages, ready for import and use in your projects.
Customizable Package Sources: Extend the package manager to support different package sources by adding to the package mapping.
Getting Started
Installation
Clone the repository to your local machine:

bash
Copy code
git clone https://github.com/yourusername/gpm.git
cd gpm
Build the GPM tool:

bash
Copy code
go build -o gpm
Make the GPM tool executable globally (optional):

bash
Copy code
# On Linux/MacOS
sudo mv gpm /usr/local/bin/gpm

# On Windows (PowerShell)
Move-Item .\gpm.exe -Destination "C:\path\to\your\preferred\location"
Usage
To use GPM, run the following command:

bash
Copy code
gpm <package-name>
Example:

bash
Copy code
gpm jwt-go
This command fetches the specified package using the Go toolchain and prepares it for use in your project.

Adding Packages
You can add custom package mappings in the packageMap variable found in the main.go file:

go
Copy code
var packageMap = map[string]string{
    "jwt-go": "github.com/dgrijalva/jwt-go",
    // Add more package mappings here if needed
}
Running the Tool
After installation, you can fetch any package listed in the mapping by simply typing:

bash
Copy code
gpm <package-name>
Video Proof of Concept (POC)
Check out the video walkthrough of the GPM tool's proof of concept (POC), where I demonstrate its basic functionality and explain the vision behind the project.

Contributing
Feedback and contributions are welcome! Feel free to fork the repository, submit issues, and create pull requests. For significant changes, please open an issue first to discuss what you would like to contribute.
