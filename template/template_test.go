package temp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"text/template"
	"time"
)

type TableEntity struct {
	Table  string
	Column map[string]string
}

func TestTemp(t *testing.T) {
	tm := template.Must(template.ParseFiles("model.tpl"))
	e := TableEntity{"testtable", map[string]string{
		"Col1": "int",
		"Col2": "string",
	}}
	tm.Execute(os.Stdout, e)
}

func MillisecondToDateString(ms int64) string {
	timeLayout := "2006-01-02"
	return time.Unix(ms/1000, 0).Format(timeLayout)
}

func TestWordBook(t *testing.T) {
	str := `{"examin":[{"Exdate":1561061258391,"Val":23,"Typ":"纪律规范"},{"Exdate":1561061258391,"Val":18,"Typ":"学习情况"},{"Exdate":1561061258391,"Val":6,"Typ":"综合素养"},{"Exdate":1561061258391,"Val":8,"Typ":"奖励加分"}],"flunk":[{"Term":1,"No":"330722199812092611","Grade":0,"Name":"必备外的技能证书学分","Passgrade":60},{"Term":1,"No":"330722199812092611","Grade":0,"Name":"各类比赛获奖学分","Passgrade":60},{"Term":1,"No":"330722199812092611","Grade":5,"Name":"音乐","Passgrade":60},{"Term":1,"No":"330722199812092611","Grade":8,"Name":"特色社会主义","Passgrade":60},{"Term":1,"No":"330722199812092611","Grade":8,"Name":"企业管理","Passgrade":60}],"interview":[{"Company":"永康威力科技股份有限公司","Post":"数控铣岗","Treatment":"额温柔温柔","Isemployed":1}],"pass":[],"plan":[{"Post":"数控铣岗","Start":1561173764000,"End":1561173766000,"Company":"永康威力科技股份有限公司"},{"Post":"钳工","Start":1559664000000,"End":1560355200000,"Company":"浙江铁牛汽车车身有限公司"},{"Post":"数控铣岗","Start":1559944089000,"End":1560441600000,"Company":"浙江铁牛汽车车身有限公司"}],"punish":[{"RecordAt":1561478400000,"Detail":"了模糊葡萄XP你也是","Rule":"讲脏话"}],"reward":[],"stu":{"enroll_year":2014,"gender":"男","mobile":"13967936015","name":"徐 最","no":"330722199812092611","org":"数控141","org_id":"aade804e5d8811e9b1ed8c1645d34af8/37fca1335ba611e9b8558c1645d34af8"}}`
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(str), &m); err != nil {
		fmt.Println(err)
	}
	tmp, err := template.New("业务档案.xml").Funcs(template.FuncMap{"dateFormat": MillisecondToDateString}).ParseFiles("业务档案.xml")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//err = tmp.Execute(w, data)
	// tm := template.Must(template.New("tmp").Funcs(template.FuncMap{"dateFormat": MillisecondToDateString}).ParseFiles("业务档案.xml"))
	f, _ := os.Create("res.xml")
	fmt.Println(tmp.Execute(f, m))
}

func TestTemplate(h *testing.T) {

	tmpl, err := template.New("tmpl").Parse("{{if .r1}}3{{else}}8{{end}}")

	if err != nil {
		panic(err)
	}
	fmt.Println(parseString(tmpl, map[string]interface{}{
		"r1":    true,
		"House": "yle",
	}))

	// buf := &bytes.Buffer{}
}

func parseString(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		panic(err)
	}
	return tmplBytes.String()
}
