package template_test

import (
	"github.com/sinlov/drone-info-tools/template"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RenderTrim(t *testing.T) {
	// mock _RenderTrim
	t.Logf("~> mock _RenderTrim")
	// do _RenderTrim
	t.Logf("~> do _RenderTrim")
	s := struct {
		Foo string
	}{
		Foo: "bar",
	}

	trim, err := template.RenderTrim("", s)
	if err != nil {
		t.Fatalf("RenderTrim error %v", err)
	}
	// verify _RenderTrim
	assert.Equal(t, "", trim)

	var tpl = `{
	"Foo": "{{ Foo }}"
}
`

	trim2, err := template.RenderTrim(tpl, s)
	if err != nil {
		t.Fatalf("RenderTrim error %v", err)
	}
	// verify _RenderTrim
	assert.Equal(t, `{
	"Foo": "bar"
}`, trim2)

	var successCheckTpl = `
{{#success Status }}✅{{/success}}
{{#failure Status}}❌{{/failure}}
`
	sModeSucc := struct {
		Status string
	}{
		Status: "success",
	}
	trim3, err := template.RenderTrim(successCheckTpl, sModeSucc)
	if err != nil {
		t.Fatalf("RenderTrim error %v", err)
	}
	sModeFail := struct {
		Status string
	}{
		Status: "failure",
	}
	assert.Equal(t, `✅`, trim3)
	trim4, err := template.RenderTrim(successCheckTpl, sModeFail)
	if err != nil {
		t.Fatalf("RenderTrim error %v", err)
	}
	assert.Equal(t, `❌`, trim4)
}
