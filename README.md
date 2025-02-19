# PappingGoLang

PappingGoLang is a simple and efficient network scanner written in **Go**. It allows users to perform **port scanning** and **network diagnostics** with ease.

## 🚀 Features

- Fast and lightweight
- Scans multiple ports simultaneously
- Supports IPv4 & IPv6
- Customizable scan ranges
- Cross-platform (Windows, Linux, macOS)

## 🛠 Installation

To build this project on your local machine, follow these steps:

1. **Clone the repository**:

   ```sh
   git clone https://github.com/Yassine3412/PappingGoLang.git
   cd PappingGoLang
   ```

2. **Build the project**:

   ```sh
   go build -o papping papping.go
   ```

3. **Run the executable**:
   ```sh
   ./papping -host 192.168.1.1 -port 80
   ```

## 📌 Usage

```
Usage: papping [options]

Options:
  -host <IP/Hostname>     Specify the target to scan (default: 0.0.0.0)
  -port <port>            Specify a single port to scan (default: 80)
```

### Example:

```sh
./papping -host scanme.nmap.org -port 443
```

## 🤝 Contributing

Pull requests are welcome! Feel free to submit issues or improvements.

## 📧 Contact

For questions or support open a ticket
