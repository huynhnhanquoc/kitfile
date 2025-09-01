
# Kit File
![Go](https://img.shields.io/badge/Go-1.21-blue)
![License](https://img.shields.io/badge/License-MIT-green)

> **[English](README.md) · [Tiếng Việt](README.vi.md)**

`kitfile` là một **gói Go nhẹ** giúp quản lý **file local** một cách dễ dàng.  

- Lấy tên file, phần mở rộng và thư mục cha  
- Thêm hoặc xóa **prefix đường dẫn**  
- Thêm hoặc xóa **hậu tố trong tên file**  
- Kiểm tra file có tồn tại và lấy đường dẫn tuyệt đối  

**Lưu ý:** Gói này chỉ hỗ trợ **file local** (không hỗ trợ URL).


## 🚀 Cài đặt

```bash
go get github.com/huynhnhanquoc/kitfile
````


## 💡 Ví dụ nhanh

Bạn có thể thử ví dụ này trực tiếp trên **Go Playground**: [Chạy trên Go Playground](https://go.dev/play/)

```go
package main

import (
    "fmt"
    "github.com/huynhnhanquoc/kitfile"
)

func main() {
    f := kitfile.New("/abc/dev.go")

    // Thêm prefix và thêm hậu tố
    f.PrependPath("xyz").AddToName(".min")
    fmt.Println(f.Location()) // /xyz/abc/dev.min.go

    // Xóa prefix và hậu tố
    f.RemovePrefixPath("xyz").RemoveFromName(".min")
    fmt.Println(f.Location()) // /abc/dev.go
}
```

**Lưu ý:** Trên Go Playground, các hàm thao tác với file hệ thống như `Exist()` hoặc `Abs()` sẽ không hoạt động, nhưng các thao tác xử lý tên và path (`PrependPath`, `RemovePrefixPath`, `AddToName`, `RemoveFromName`) vẫn chạy được.


## 📚 Tổng Quan API

| Hàm                                     | Mô tả                                       |
| --------------------------------------- | ------------------------------------------- |
| `New(location string) *File`            | Tạo một instance File mới                   |
| `Exist() error`                         | Kiểm tra file có tồn tại không              |
| `Name() string`                         | Lấy tên file (basename)                     |
| `NameWithoutExt() string`               | Lấy tên file bỏ phần mở rộng                |
| `Ext() string`                          | Lấy phần mở rộng của file                   |
| `Dir() string`                          | Lấy thư mục cha của file                    |
| `Abs() (string, error)`                 | Lấy đường dẫn tuyệt đối                     |
| `Location() string`                     | Lấy đường dẫn nguyên gốc                    |
| `PrependPath(prefix string) *File`      | Thêm prefix vào trước đường dẫn hiện tại    |
| `RemovePrefixPath(prefix string) *File` | Xóa prefix ở đầu đường dẫn nếu có           |
| `AddToName(suffix string) *File`        | Thêm hậu tố trước phần mở rộng của tên file |
| `RemoveFromName(suffix string) *File`   | Xóa hậu tố trước phần mở rộng nếu có        |


## 👤 Tác giả

**Huỳnh Nhân Quốc** – [huynhnhanquoc.com](https://huynhnhanquoc.com)

GitHub: [github.com/huynhnhanquoc](https://github.com/huynhnhanquoc)


## ☕ Hỗ trợ tôi

Nếu bạn thấy gói này hữu ích, hãy cân nhắc **hỗ trợ tôi trên Buy Me a Coffee**:

[![Buy Me a Coffee](https://img.shields.io/badge/Buy%20Me%20a%20Coffee-☕-ff813f)](https://www.buymeacoffee.com/huynhnhanquoc)


## 📄 Giấy phép

MIT License © Huỳnh Nhân Quốc

