package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	t.Run("should return textnode string data", func(t *testing.T) {
		txt := []byte("<section><h1>Hello Parser</h1></section>")
		pt := Parse(txt)

		tx := Text("Hello Parser")
		assert.Equal(t, tx.Data, pt.Children[0].Children[0].Data)
	})
}

func Test_ParseText(t *testing.T) {
	t.Run("should create a valid textnode", func(t *testing.T) {
		p := Parser{pos: 4, input: []byte("<h1>Hello Parser</h1>")}
		txN := p.ParseText()
		tn := Text("Hello Parser")

		assert.Equal(t, tn.Data, txN.Data)
	})
}

func Test_ParseElement(t *testing.T) {
	p := Parser{pos: 0, input: []byte("<h1>Hello Parser</h1>")}
	en, err := p.ParseElement()
	if err != nil {
		t.Error()
		return
	}
	t.Run("tagName should be 'h1'", func(t *testing.T) {
		assert.Equal(t, "h1", en.TagName)
	})

	t.Run("should have one child", func(t *testing.T) {
		tn := Node{NodeType: TextNode, Data: "Hello Parser", Children: []Node{}}
		assert.Equal(t, []Node{tn}, en.Children)
	})

	t.Run("should return error for invalid HTMLElement", func(t *testing.T) {
		p := Parser{pos: 0, input: []byte("h1>Hello Parser</h1>")}
		_, err := p.ParseElement()
		if err == nil {
			t.Errorf("%s", "should return error for invalid html element")
		}
		assert.EqualError(t, err, "invalid HTMLElement: opening tag not found")
	})
}
