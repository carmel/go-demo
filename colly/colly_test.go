package colly

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"github.com/jmoiron/sqlx"
)

type (
	course struct {
		Jxbzc string
		KchId string `json:"kch_id"`
		Xqj   int32
		Jxbmc string
		Jxbrs int32
		Jc    int32
		Sksj  string
		Zcd   int32
		JxbId string `json:"jxb_id"`
		Kcmc  string
		Jxdd  string
		Rq    string
	}
	performance struct {
		XhId   string `json:"xh_id"`
		BhId   string `json:"bh_id"`
		Njmc   string `json:"njmc"`
		Xbm    string `json:"xbm"`
		Jgmc   string `json:"jgmc"`
		Xb     string `json:"xb"`
		Zymc   string `json:"zymc"`
		JxdmId string `json:"jxdm_id"`
		Xh     string `json:"xh"`
		Bjmc   string `json:"bjmc"`
		Qdfs   string `json:"qdfs"`
		Xm     string `json:"xm"`
		NjdmId string `json:"njdm_id"`
		ZyhId  string `json:"zyh_id"`
		JgId   string `json:"jg_id"`
		Dmlbm  string `json:"dmlbm"`
	}
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	var err error
	db, err = sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/student?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	m.Run()
}

// 获取财经大学学生考勤信息
func TestAttendanceInfo(t *testing.T) {
	defer db.Close()
	// start scraping
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: jwxt.shufe-zj.edu.cn
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"),
		colly.AllowedDomains("jwxt.shufe-zj.edu.cn", "ty.shufe-zj.edu.cn"),
	)

	// Set max Parallelism and introduce a Random Delay
	// c.Limit(&colly.LimitRule{
	// 	Parallelism: 2,
	// 	RandomDelay: 5 * time.Second,
	// })

	// Create another collector to scrape details
	dc := c.Clone()

	var rq string

	dc.OnResponse(func(r *colly.Response) {
		if r.Body != nil {
			var pm []map[string]interface{}
			if err := json.Unmarshal(r.Body, &pm); err != nil {
				t.Fatal("performance parse failed", err)
			}
			for _, v := range pm {
				v["rq"] = rq
				fmt.Println(SaveMap("performance", v))
			}
			// fmt.Printf("%+v\n", pm)
		}
	})

	dc.OnRequest(func(r *colly.Request) {
		fmt.Println("访问: ", r.URL.String())
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		// r.Headers.Set("Referrer", "http://jwxt.shufe-zj.edu.cn/jwglxt/jxdmgl/jsjxdm_cxJsjxdmIndex.html?gnmkdm=N254350&layout=default&su=021560")
		r.Headers.Set("Cookie", "JSESSIONID=43AE3C714A6A2D4500888D7D4CBE1258; route=365670b7b823d6a3440639605a12dfa9")
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("访问: ", r.URL.String())
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		// r.Headers.Set("Referrer", "http://jwxt.shufe-zj.edu.cn/jwglxt/jxdmgl/jsjxdm_cxJsjxdmIndex.html?gnmkdm=N254350&layout=default&su=021560")
		r.Headers.Set("Cookie", "JSESSIONID=43AE3C714A6A2D4500888D7D4CBE1258; route=365670b7b823d6a3440639605a12dfa9")
	})

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		if r.Body != nil {
			var cs []map[string]interface{}
			if err := json.Unmarshal(r.Body, &cs); err != nil {
				t.Fatal("course parse failed", err)
			}

			// fmt.Printf("%+v\n", cs)
			time.Sleep(time.Duration(2) * time.Second)
			for _, v := range cs {
				fmt.Println(SaveMap("course", v))
				rq = v["rq"].(string)
				if err := dc.Post("http://jwxt.shufe-zj.edu.cn/jwglxt/jxdmgl/jsjxdm_cxJxbxxByJxbid.html?gnmkdm=N254350&su=021560", map[string]string{
					"jxb_id": v["jxb_id"].(string),
					"rq":     rq,
				}); err != nil {
					t.Fatal(err)
				}
				time.Sleep(time.Duration(2) * time.Second)
			}
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// 格式化
	const layout = "2006-01-02 15:04:05" //时间常量
	Loc, _ := time.LoadLocation("Asia/Shanghai")
	cDate, _ := time.ParseInLocation(layout, "2020-05-12 00:00:00", Loc)

	for cDate.Format("2006-01-02") != "2020-06-12" {
		if err := c.Post("http://jwxt.shufe-zj.edu.cn/jwglxt/jxdmgl/jsjxdm_cxKcListByRq.html?gnmkdm=N254350&su=021560", map[string]string{
			"rq": cDate.Format("2006-01-02"),
		}); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Duration(2) * time.Second)
		cDate = cDate.Add(86400000000000)
	}
}

func TestT(t *testing.T) {
	var x struct {
		A string `json:"atag"`
		B int    `json:"btag"`
	}
	// s := reflect.TypeOf(x) //通过反射获取type定义
	// r := regexp.MustCompile(`json:"(\w+)"`)
	// for i := 0; i < s.NumField(); i++ {
	// 	fmt.Println(r.FindStringSubmatch(string(s.Field(i).Tag))[1], s.Field(i).Type) //将tag输出出来
	// }
	var p performance
	var c course
	structToTable(x, p, c)
}

func structToTable(structs ...interface{}) {
	var sb strings.Builder
	var tag string
	r := regexp.MustCompile(`json:"(\w+)"`)
	for _, sct := range structs {
		s := reflect.TypeOf(sct)
		sb.WriteString(`CREATE TABLE `)
		sb.WriteString(s.Name())
		sb.WriteString(" (\n")
		fn := s.NumField()
		for i := 0; i < fn; i++ {
			sb.WriteString("  ")
			if tag = string(s.Field(i).Tag); tag == "" {
				sb.WriteString(s.Field(i).Name)
			} else {
				sb.WriteString(r.FindStringSubmatch(tag)[1])
			}
			sb.WriteString("	")
			switch s.Field(i).Type.Name() {
			case "int", "int16", "int32":
				sb.WriteString("INT")
			case "int64":
				sb.WriteString("BIGINT")
			case "string":
				fallthrough
			default:
				sb.WriteString("VARCHAR(50)")
			}

			if i+1 != fn {
				sb.WriteString(",")
			}
			sb.WriteString("\n")
		}
		sb.WriteString(");\n\n")
	}
	fmt.Println(sb.String())
}

func SaveMap(t string, m map[string]interface{}) error {
	var err error
	var add []string
	var edit []string
	for f, v := range m {
		var flag bool
		switch z := v.(type) { // 零值过滤
		case float32: // 注意float32与float64不可写在一起，因在case路由中，如果不能精准到单路线，v还是一个interface{}
			flag = math.Abs(float64(z)-0) < 0.0000001
		case float64:
			flag = math.Abs(z-0) < 0.0000001
		case int, int16, int32, int64:
			flag = v == 0
		case string:
			flag = v == ""
		case nil:
			flag = true
		}
		if flag {
			continue
		}
		add = append(add, f)
		if f != "Id" && f != "id" {
			edit = append(edit, f+"=:"+f)
		}
	}

	if len(m) != len(edit) {
		_, err = db.NamedExec(`UPDATE `+t+` SET `+strings.Join(edit, ",")+` WHERE id = :id`, m)
	} else {
		// xid.New().String()  // 假如要传ID
		_, err = db.NamedExec(`INSERT INTO `+t+`(`+strings.Join(add, ",")+`)VALUES(`+`:`+strings.Join(add, ",:")+`)`, m)
	}

	return err
}
