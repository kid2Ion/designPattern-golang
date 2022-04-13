package adapter

// 再利用性(Embedded)(委譲)

//target
type Decorator interface {
	Decorate() string
}

//目的
type Banner struct {
	str string
}

// 目的
func (self *Banner) getString() string {
	return "*" + self.str + "*"
}

//adaptor
type EmbeddedDecorateBanner struct {
	*Banner
}

//インスタンス化(埋め込みでも値setの際は明治的に示すので、ラップして呼び出し側が意識しないでいいようにする)
func NewEmbeddedDecorateBanner(str string) *EmbeddedDecorateBanner {
	return &EmbeddedDecorateBanner{&Banner{str}}
}

func (self *EmbeddedDecorateBanner) Decorate() string {
	return self.getString()
}
