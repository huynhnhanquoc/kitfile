# Kit File

![Go](https://img.shields.io/badge/Go-1.21-blue)
![License](https://img.shields.io/badge/License-MIT-green)

> **[English](README.md) · [Tiếng Việt](README.vi.md)**

`kitfile` là một **thư viện Go nhẹ** giúp quản lý **tệp cục bộ** một cách dễ dàng.

* Lấy tên file, phần mở rộng, và thư mục cha
* Thêm hoặc xoá **tiền tố đường dẫn**
* Thêm hoặc xoá **hậu tố trong tên file**
* Đọc, ghi và ghi an toàn vào file với khả năng tự động tạo thư mục
* Kiểm tra sự tồn tại của file và lấy đường dẫn tuyệt đối

**Lưu ý:** Thư viện này chỉ hỗ trợ **file cục bộ** (không hỗ trợ URL).


## 🚀 Cài đặt

```bash
go get github.com/huynhnhanquoc/kitfile
```


## 💡 Ví dụ nhanh

Bạn có thể chạy trực tiếp ví dụ này trong **Go Playground**: [Chạy trên Go Playground](https://go.dev/play/)

```go
package main

import (
    "fmt"
    "github.com/huynhnhanquoc/kitfile"
)

func main() {
    f := kitfile.New("/abc/dev.go")

    // Thêm đường dẫn và gắn hậu tố
    f.PrependPath("xyz").AddToName(".min")
    fmt.Println(f.Location()) // /xyz/abc/dev.min.go

    // Xoá tiền tố và hậu tố
    f.RemovePrefixPath("xyz").RemoveFromName(".min")
    fmt.Println(f.Location()) // /abc/dev.go

    // Ghi vào file (tự động tạo thư mục nếu chưa có)
    err := f.Write([]byte("hello world"))
    if err != nil {
        panic(err)
    }

    // Đọc file
    content, _ := f.Read()
    fmt.Println(string(content)) // hello world
}
```


## 📚 Tổng quan API

| Hàm                                              | Mô tả                                                           |
| ------------------------------------------------ | --------------------------------------------------------------- |
| `New(location string) *File`                     | Tạo một đối tượng File mới                                      |
| `NewSafe(location string) (*File, error)`        | Tạo đối tượng File đã được kiểm tra hợp lệ (làm sạch đường dẫn) |
| `Exist() error`                                  | Kiểm tra file có tồn tại hay không                              |
| `Name() string`                                  | Lấy tên file (basename)                                         |
| `NameWithoutExt() string`                        | Lấy tên file không có phần mở rộng                              |
| `Ext() string`                                   | Lấy phần mở rộng của file                                       |
| `Dir() string`                                   | Lấy thư mục cha                                                 |
| `Abs() (string, error)`                          | Lấy đường dẫn tuyệt đối                                         |
| `Location() string`                              | Lấy đường dẫn gốc của file                                      |
| `Read() ([]byte, error)`                         | Đọc nội dung file dưới dạng byte                                |
| `WriteSafe(data []byte) error`                   | Ghi dữ liệu vào **file mới** (lỗi nếu file đã tồn tại)          |
| `Write(data []byte) error`                       | Ghi dữ liệu (tạo mới hoặc ghi đè, tự động tạo thư mục)          |
| `WritePermission(data []byte, perm os.FileMode)` | Ghi dữ liệu với quyền tùy chỉnh (tự động tạo thư mục)           |
| `PrependPath(prefix string) *File`               | Thêm tiền tố vào đường dẫn hiện tại                             |
| `RemovePrefixPath(prefix string) *File`          | Xoá tiền tố trong đường dẫn nếu có                              |
| `AddToName(suffix string) *File`                 | Thêm hậu tố trước phần mở rộng                                  |
| `RemoveFromName(suffix string) *File`            | Xoá hậu tố trong tên file nếu có                                |


## 👤 Tác giả

**Huỳnh Nhân Quốc** – [huynhnhanquoc.com](https://huynhnhanquoc.com)

GitHub: [github.com/huynhnhanquoc](https://github.com/huynhnhanquoc)


## ☕ Ủng hộ tôi

Nếu bạn thấy dự án này hữu ích, hãy cân nhắc ủng hộ tôi qua **Buy Me a Coffee**:

[![Buy Me a Coffee](https://img.shields.io/badge/Buy%20Me%20a%20Coffee-☕-ff813f)](https://www.buymeacoffee.com/huynhnhanquoc)


## 📄 Giấy phép

2025 © Huỳnh Nhân Quốc - Founder of [Kit Module](https://kitmodule.com)

Phát hành theo [Giấy phép MIT](https://github.com/huynhnhanquoc/kitfile/blob/master/LICENSE)
