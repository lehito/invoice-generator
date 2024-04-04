package invoice

import (
	"os"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestGenerateInvoice(t *testing.T) {
	builder, err := NewBuilder(
		Config{
			SealImage:      "./sample-seal.png",
			FontName:       "noto-sans-cjk",
			FontNormal:     "./fonts/NotoSansCJK-JP/NotoSansCJKjp-Regular.ttf",
			FontItalic:     "./fonts/NotoSansCJK-JP/NotoSansCJKjp-Italic.ttf",
			FontBold:       "./fonts/NotoSansCJK-JP/NotoSansCJKjp-Bold.ttf",
			FontBoldItalic: "./fonts/NotoSansCJK-JP/NotoSansCJKjp-BoldItalic.ttf",
		},
		"./sample-params.yaml")
	if err != nil {
		t.Fatal("failed to create builder")
		return
	}

	buf, err := builder.GenerateInvoice()
	if buf == nil || err != nil {
		t.Fatal("failed to generate invoice")
		return
	}

	filename := "sample-invoice.pdf"
	if err := os.WriteFile(filename, buf, 0666); err != nil {
		t.Fatal("failed to write to file")
		return
	}
}
