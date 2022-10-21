package gooxml

import (
	"testing"

	"github.com/carmel/gooxml/color"
	"github.com/carmel/gooxml/common"
	"github.com/carmel/gooxml/document"
	"github.com/carmel/gooxml/measurement"
	"github.com/carmel/gooxml/schema/soo/wml"
)

func TestArchive(t *testing.T) {
	doc := document.New()

	img1, _ := common.ImageFromFile("1.png")
	img1ref, _ := doc.AddImage(img1)

	para := doc.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run := para.AddRun()
	// run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	para.SetStyle("Title")
	run.AddText("学徒业务档案")

	table := doc.AddTable()
	// 4 inches wide
	table.Properties().SetWidthPercent(100)
	table.Properties().Borders().SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)
	table.Properties().SetAlignment(wml.ST_JcTableCenter)

	row := table.AddRow()
	row.Properties().SetHeight(0.4*measurement.Inch, wml.ST_HeightRuleExact)

	cell := row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(3)
	cell.Properties().SetWidthPercent(26)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("姓名")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	cell.Properties().SetWidthPercent(54)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("张三")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
	cell.Properties().SetColumnSpan(2)
	cell.Properties().SetWidthPercent(20)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	para.AddRun().AddDrawingInline(img1ref)
	///////////////////////////////////////////
	row = table.AddRow()
	row.Properties().SetHeight(0.4*measurement.Inch, wml.ST_HeightRuleExact)

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(3)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("班级")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("33班")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	cell.Properties().SetColumnSpan(2)
	para = cell.AddParagraph()
	///////////////////////////////////////////
	row = table.AddRow()
	row.Properties().SetHeight(0.4*measurement.Inch, wml.ST_HeightRuleExact)

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(3)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("身份证")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("421127")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	cell.Properties().SetColumnSpan(2)
	para = cell.AddParagraph()
	///////////////////////////////////////////
	row = table.AddRow()
	row.Properties().SetHeight(0.4*measurement.Inch, wml.ST_HeightRuleExact)

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(3)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("手机号")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("138")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	cell.Properties().SetColumnSpan(2)
	para = cell.AddParagraph()
	///////////////////////////////////////////

	row = table.AddRow()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
	cell.Properties().SetWidthAuto()
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	// run.Properties().SetVerticalAlignment(sharedTypes.ST_VerticalAlignRunBaseline)
	run.AddText("学业课程")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(2)
	cell.Properties().SetWidthAuto()
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("优秀课程")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(6)
	cell.Properties().SetWidthAuto()
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("优秀课程1，优秀课程2")

	row = table.AddRow()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	// cell.Properties().SetWidthPercent(10)
	para = cell.AddParagraph()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(2)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("不及格课程")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(6)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("不及格课程1，不及格课程2")

	// 德语
	row = table.AddRow()
	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
	cell.Properties().SetWidthAuto()
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	// run.Properties().SetVerticalAlignment(sharedTypes.)
	run.AddText("德育学分")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	// cell.Properties().SetWidth(1.3 * measurement.Inch)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("第一学期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	// cell.Properties().SetWidth(0.8 * measurement.Inch)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("得分")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetWidthAuto()
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("第二学期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetWidthAuto()
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("得分")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetWidthAuto()
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("第三学期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetWidthAuto()
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("得分")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetWidthAuto()
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("第四学期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetWidthAuto()
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("得分")

	row = table.AddRow()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	// cell.Properties().SetWidthPercent(10)
	para = cell.AddParagraph()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("第五学期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("得分")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("第六学期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("得分")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("第七学期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("得分")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("第八学期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("得分")

	row = table.AddRow()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	para = cell.AddParagraph()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("第九学期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("得分")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("第十学期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("得分")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("第十一学期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("得分")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("第十二学期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("得分")

	// 技能证书
	row = table.AddRow()
	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	// run.Properties().SetVerticalAlignment(sharedTypes.)
	run.AddText("技能证书")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("证书名称")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("等级")

	row = table.AddRow()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	para = cell.AddParagraph()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("证书1")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("等级A")

	// 奖励记录
	row = table.AddRow()
	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	// run.Properties().SetVerticalAlignment(sharedTypes.)
	run.AddText("奖励记录")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("奖励项目")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("等级")

	row = table.AddRow()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	para = cell.AddParagraph()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("奖励1")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("等级A")

	// 处罚记录
	row = table.AddRow()
	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	// run.Properties().SetVerticalAlignment(sharedTypes.)
	run.AddText("处罚记录")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("处罚项目")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("等级")

	row = table.AddRow()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	para = cell.AddParagraph()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("处罚1")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(4)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("等级A")

	// 实践经历
	row = table.AddRow()
	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	// run.Properties().SetVerticalAlignment(sharedTypes.)
	run.AddText("实践经历")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("开始日期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("截至日期")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("等级")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(2)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("岗位")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(3)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("单位")

	row = table.AddRow()
	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	para = cell.AddParagraph()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("2019-10-1")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("2020-10-1")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("优秀")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(2)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("磨具磨料")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(3)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("铁牛")

	// 面试经历
	row = table.AddRow()
	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	// run.Properties().SetVerticalAlignment(sharedTypes.)
	run.AddText("面试经历")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("是否录用")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(2)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("待遇情况")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(2)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("岗位")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(3)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("单位")

	row = table.AddRow()
	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	para = cell.AddParagraph()

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("已录用")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(2)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("400/月")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(2)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("焊工")

	cell = row.AddCell()
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.Properties().SetColumnSpan(3)
	para = cell.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetFontFamily("仿宋")
	run.Properties().SetSize(12)
	run.AddText("铁牛")

	doc.SaveToFile("arc.docx")
}
