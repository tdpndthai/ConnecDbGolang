package computer

type Spec struct{
	Maker string //vì viết chữ hoa nên có thể truy nhập được từ file khác có cùng package
	model string //không thể  truy nhập được
	Price int
}