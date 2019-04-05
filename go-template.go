package main
import(
	"text/template"
	"os"
)
type TableEntity struct{
	Table string
	Column map[string]string
}
func main(){
	t := template.Must(template.ParseFiles("template/model.tpl"))
	e := TableEntity{"testtable",map[string]string{
		"Col1":"int",
		"Col2":"string",
	}}
	t.Execute(os.Stdout,e)
}