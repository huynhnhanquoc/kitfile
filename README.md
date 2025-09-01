
# Kit File
![Go](https://img.shields.io/badge/Go-1.21-blue)
![License](https://img.shields.io/badge/License-MIT-green)

> **[English](README.md) · [Tiếng Việt](README.vi.md)**

`kitfile` is a **lightweight Go package** for managing **local files** with ease.  

- Retrieve file name, extension, and parent directory  
- Prepend or remove **prefix paths**  
- Add or remove **suffixes in file names**  
- Check if a file exists and get absolute paths  

**Note:** This package only supports **local files** (no URLs).


## 🚀 Installation

```bash
go get github.com/huynhnhanquoc/kitfile
````


## 💡 Quick Example
You can try this example directly in **Go Playground**: [Run on Go Playground](https://go.dev/play/)

```go
package main

import (
    "fmt"
    "github.com/huynhnhanquoc/kitfile"
)

func main() {
    f := kitfile.New("/abc/dev.go")

    // Prepend a path and add a suffix
    f.PrependPath("xyz").AddToName(".min")
    fmt.Println(f.Location()) // /xyz/abc/dev.min.go

    // Remove prefix and suffix
    f.RemovePrefixPath("xyz").RemoveFromName(".min")
    fmt.Println(f.Location()) // /abc/dev.go
}
```


## 📚 API Overview

| Function                                | Description                                          |
| --------------------------------------- | ---------------------------------------------------- |
| `New(location string) *File`            | Create a new File instance                           |
| `Exist() error`                         | Check if the file exists                             |
| `Name() string`                         | Get the file name (basename)                         |
| `NameWithoutExt() string`               | Get the file name without extension                  |
| `Ext() string`                          | Get the file extension                               |
| `Dir() string`                          | Get the parent directory                             |
| `Abs() (string, error)`                 | Get the absolute path                                |
| `Location() string`                     | Get the raw file path                                |
| `PrependPath(prefix string) *File`      | Add a prefix path before the current path            |
| `RemovePrefixPath(prefix string) *File` | Remove a prefix path if present                      |
| `AddToName(suffix string) *File`        | Add a suffix before the file extension               |
| `RemoveFromName(suffix string) *File`   | Remove a suffix before the file extension if present |


## 👤 Author

**Huỳnh Nhân Quốc** – [huynhnhanquoc.com](https://huynhnhanquoc.com)

GitHub: [github.com/huynhnhanquoc](https://github.com/huynhnhanquoc)


## ☕ Support Me

If you find this project useful, consider supporting me on **Buy Me a Coffee**:

[![Buy Me a Coffee](https://img.shields.io/badge/Buy%20Me%20a%20Coffee-☕-ff813f)](https://www.buymeacoffee.com/huynhnhanquoc)


## 📄 License
2025 © Huỳnh Nhân Quốc - Founder of [Kit Module](https://kitmodule.com) 

Released under the [MIT License](https://github.com/huynhnhanquoc/kitfile/blob/master/LICENSE)
