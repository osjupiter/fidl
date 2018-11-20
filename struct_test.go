package dfl

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

var struct1 Record
var struct2 Record
var struct3 Record
var graph Graph

func init() {
	g := Graph{
		records: []Record{
			Record{
				name: "struct1",
				values: []Value{
					{"left", "f0"},
					{"middle", "f1"},
					{"right", "f2"},
				},
			},
			Record{
				name: "struct2",
				values: []Value{
					{"left", "f0"},
					{"middle", "f1"},
					{"right", "f2"},
				},
			},
			Record{
				name: "struct3",
				values: []Value{
					{"left", "f0"},
					{"middle", "f1"},
					{"right", "f2"},
				},
			},
		},
		arrows: []Arrow{
			{src: "struct1:f1", dest: "struct2:f0"},
			{src: "struct2:f2", dest: "struct3:f1"},
			{src: "struct2:f2", dest: "struct3:f2"},
		},
		template:SimpleTempate{},
	}
	graph = g
}

// グラフ全体を表すオブジェクトの出力をサンプルdotファイルと比較する
func TestOutput(t *testing.T) {
	b, e := ioutil.ReadFile("test/test.dot")
	if e != nil {
		t.Fatal(e)
	}
	ans := string(b)
	if graph.String() != ans {
		b:=bytes.Buffer{}
		b.WriteString(graph.String())
		ioutil.WriteFile("hoge",b.Bytes() ,0777)
		t.Fatalf("invalid answer:" + graph.String())
	}

}

// record表現オブジェクトの出力を確認する
func TestRecordExpression(t *testing.T) {
	r := Record{name: "struct1",
		values: []Value{
			{"left", "f0"},
			{"middle", "f1"},
			{"right", "f2"},
		}}
	answer := `[label="struct1|<f0> left|<f1> middle|<f2> right"]`
	if r.String() != answer {
		t.Fatalf("invalid answer:" + r.String())
	}

}

// value表現オブジェクトの出力を確認する
func TestValueExpression(t *testing.T) {
	v := Value{name: "testname", port: "testport"}
	answer := "<testport> testname"
	if v.String() != answer {
		t.Fatalf("invalid answer:" + v.String())
	}
}

// arrow表現オブジェクトの出力を確認する
func TestArrowExpression(t *testing.T) {
	a := Arrow{
		src:  "struct1:f1",
		dest: "struct2:f0",
	}
	answer := "struct1:f1 -> struct2:f0"
	if a.String() != answer {
		t.Fatalf("invalid answer:" + a.String())
	}
}

// record,arrowによらないファイルのテンプレート出力内容を確認する
func TestHeaderMisc(t *testing.T) {
	str := SimpleTempate{}.head()
	if !strings.Contains(str, "rankdir=TB") {
		t.Fatalf("template error")
	}
}
