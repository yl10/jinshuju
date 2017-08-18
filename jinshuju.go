package jinshuju

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

//JSJ 金数据
type JSJ struct {
	Key    string
	Secret string
}

//Form 金数据表单
type Form struct {
	Name        string
	Description string
	Fields      []map[string]Field
}

const (
	//SingleChoiceField 单选
	SingleChoiceField FieldType = iota + 1
	//MultipleChoiceField 多选
	MultipleChoiceField

	DropDownField
	DateField
	EmailField
	NumberField
	PhoneField

	//ParagraphTextField 多行文本
	ParagraphTextField
	//SingleLineTextField 单行文本
	SingleLineTextField
)

//FieldType 字段类型
type FieldType int
type fieldtype struct {
	Name        string
	Description string
}

var fieldtypemap = map[FieldType]fieldtype{
	SingleChoiceField:   {"single_choice", "单选框"},
	MultipleChoiceField: {"multiple_choice", "复选框"},
	DropDownField:       {"drop_down", "下拉框"},
	DateField:           {"date", "日期"},
	EmailField:          {"email", "Email"},
	NumberField:         {"number", "数字"},
	PhoneField:          {"phone", "电话"},
	ParagraphTextField:  {"paragraph_text", "多行文本"},
	SingleLineTextField: {"single_line_text", "单行文本"},
}

func (f FieldType) String() string {
	if v, ok := fieldtypemap[f]; ok {
		return v.Description
	}
	return ""
}

//Name 返回类型名称
func (f FieldType) Name() string {
	if v, ok := fieldtypemap[f]; ok {
		return v.Name
	}
	return ""
}

//Field 字段
type Field struct {
	Label      string
	Type       string
	Private    bool
	Validation Validation
	Choices    []Choice
}

//Validation 校验
type Validation struct {
	Required bool
}

//Choice 选项
type Choice struct {
	Name  string
	Value string
}

func NewJSJ(key string, secret string) *JSJ {
	return &JSJ{
		key, secret,
	}
}

//GetFormInfo 获取表单定义
func (j JSJ) GetFormInfo(formid string) (*Form, error) {
	uri := "https://jinshuju.net/api/v1/forms/" + formid
	data, err := j.sendGetrequest(uri)
	if err == nil {
		var form Form
		err = json.Unmarshal(data, &form)
		return &form, err
	}
	return nil, err

}

func (j JSJ) sendGetrequest(uri string) ([]byte, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	u.User = url.UserPassword(j.Key, j.Secret)
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
