# Checklist tuân thủ quy tắc viết test trong Go

Dưới đây là danh mục các quy tắc bắt buộc phải tuân thủ khi viết mã kiểm thử (unit test) bằng công cụ tích hợp sẵn của ngôn ngữ Go, dựa trên tài liệu [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world).

## 1. Quy tắc về tệp tin và hàm kiểm thử
* [ ] **Hậu tố tên file:** Tên file chứa mã kiểm thử bắt buộc phải kết thúc bằng cụm `_test.go` (Ví dụ: `hello_test.go`).
* [ ] **Tiền tố tên hàm:** Tên của hàm kiểm thử phải bắt đầu bằng chữ `Test` viết hoa (Ví dụ: `TestHello`).
* [ ] **Tham số đầu vào:** Hàm kiểm thử chỉ nhận duy nhất một tham số là `t *testing.T`.
* [ ] **Import thư viện:** Phải khai báo `import "testing"` ở đầu file để có thể sử dụng kiểu dữ liệu `*testing.T`.

## 2. Quy tắc viết code và định dạng trong test
* [ ] **Sử dụng Helper (nếu có):** Nếu viết hàm bổ trợ để so sánh kết quả (`assertion helper`), bắt buộc phải gọi `t.Helper()` ở đầu hàm để báo lỗi hiển thị đúng dòng code bị sai ở vị trí gọi chứ không phải ở hàm helper.
* [ ] **Định dạng thông báo lỗi:** Sử dụng `%q` trong hàm `t.Errorf` để bao bọc các giá trị kết quả trong dấu ngoặc kép, giúp dễ dàng phát hiện khoảng trắng thừa hoặc chuỗi rỗng.
* [ ] **Rút gọn tham số hàm helper:** Khi truyền nhiều tham số cùng kiểu dữ liệu vào hàm helper, hãy gom lại ngắn gọn (Ví dụ: dùng `(got, want string)` thay vì `(got string, want string)`).

## 3. Quy trình thực hiện (Chu kỳ TDD)
* [ ] **Viết test trước (Test-driven):** Luôn luôn viết mã kiểm thử trước khi viết mã nguồn xử lý tính năng chính.
* [ ] **Kiểm tra lỗi biên dịch:** Chạy thử để compiler báo lỗi, sau đó sửa code vừa đủ để hệ thống biên dịch thành công.
* [ ] **Xác nhận test lỗi:** Đảm bảo bài test phải chạy thất bại với một thông báo lỗi rõ ràng, có ý nghĩa trước khi viết code xử lý logic.
* [ ] **Viết mã xử lý vừa đủ:** Chỉ viết lượng mã nguồn vừa đủ để vượt qua bài test (làm cho test chuyển sang màu xanh).
* [ ] **Tối ưu hóa (Refactor):** Tiến hành tối ưu và làm sạch cả mã nguồn chính lẫn mã test sau khi các bài test đã chạy thành công.