package view

func Hello() string {
	return `
👋 Hi! Tôi là thằng nhỏ của Bình

📌 Các lệnh bạn có thể dùng:

/hello
→ Hiển thị danh sách lệnh

/track <tracking_code>
→ Theo dõi đơn hàng
Ví dụ: /track SPXVN123456789

/untrack <tracking_code>
→ Ngừng theo dõi đơn

/myorders
→ Danh sách đơn đang theo dõi

/help
→ Hiển thị trợ giúp
`
}
