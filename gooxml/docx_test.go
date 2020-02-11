package gooxml

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/carmel/gooxml/color"
	"github.com/carmel/gooxml/common"
	"github.com/carmel/gooxml/document"
	"github.com/carmel/gooxml/measurement"
	"github.com/carmel/gooxml/schema/soo/wml"
)

var lorem = "我是一段很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长的文本"

func TestImage(t *testing.T) {
	doc := document.New()
	img1, _ := common.ImageFromFile("1.png")
	img2, _ := common.ImageFromFile("2.png")
	img3, _ := common.ImageFromFile("3.png")

	img1ref, _ := doc.AddImage(img1)
	img2ref, _ := doc.AddImage(img2)
	img3ref, _ := doc.AddImage(img3)

	{
		table := doc.AddTable()
		// 4 inches wide
		table.Properties().SetWidthPercent(100)
		table.Properties().Borders().SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)
		table.Properties().SetAlignment(wml.ST_JcTableCenter)

		row := table.AddRow()
		// row.Properties().SetHeight(2*measurement.Inch, wml.ST_HeightRuleExact)

		cell := row.AddCell()
		cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
		cell.Properties().SetColumnSpan(2)

		para := cell.AddParagraph()
		para.Properties().SetAlignment(wml.ST_JcCenter)
		run := para.AddRun()
		run.Properties().SetFontFamily("仿宋")
		run.AddText("Cells can span multiple columns")

		row = table.AddRow()
		cell = row.AddCell()
		cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
		cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
		cell.AddParagraph().AddRun().AddText("Vertical Merge")
		para = row.AddCell().AddParagraph()
		para.Properties().SetAlignment(wml.ST_JcCenter)
		para.AddRun().AddDrawingInline(img1ref)

		row = table.AddRow()
		cell = row.AddCell()
		cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
		cell.AddParagraph()
		row.AddCell().AddParagraph().AddRun().AddText("1122")

		row = table.AddRow()
		row.AddCell().AddParagraph().AddRun().AddText("Street Address")
		row.AddCell().AddParagraph().AddRun().AddText("111 Country Road")
	}
	doc.AddParagraph()
	{
		para := doc.AddParagraph()
		anchored, err := para.AddRun().AddDrawingAnchored(img1ref)
		if err != nil {
			log.Fatalf("unable to add anchored image: %s", err)
		}
		anchored.SetName("Gopher")
		anchored.SetSize(2*measurement.Inch, 2*measurement.Inch)
		anchored.SetOrigin(wml.WdST_RelFromHPage, wml.WdST_RelFromVTopMargin)
		anchored.SetHAlignment(wml.WdST_AlignHCenter)
		anchored.SetYOffset(3 * measurement.Inch)
		anchored.SetTextWrapSquare(wml.WdST_WrapTextBothSides)

		run := para.AddRun()
		for i := 0; i < 16; i++ {
			run.AddText(lorem)

			// drop an inline image in
			if i == 13 {
				inl, err := run.AddDrawingInline(img2ref)
				if err != nil {
					log.Fatalf("unable to add inline image: %s", err)
				}
				inl.SetSize(1*measurement.Inch, 1*measurement.Inch)
			}
			if i == 15 {
				inl, err := run.AddDrawingInline(img3ref)
				if err != nil {
					log.Fatalf("unable to add inline image: %s", err)
				}
				inl.SetSize(1*measurement.Inch, 1*measurement.Inch)
			}
		}
	}
	doc.SaveToFile("image.docx")
}

func TestDocproperties(t *testing.T) {
	doc, err := document.Open("document.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	cp := doc.CoreProperties
	// You can read properties from the document
	fmt.Println("Title:", cp.Title())
	fmt.Println("Author:", cp.Author())
	fmt.Println("Description:", cp.Description())
	fmt.Println("Last Modified By:", cp.LastModifiedBy())
	fmt.Println("Category:", cp.Category())
	fmt.Println("Content Status:", cp.ContentStatus())
	fmt.Println("Created:", cp.Created())
	fmt.Println("Modified:", cp.Modified())

	// And change them as well
	cp.SetTitle("CP Invoices")
	cp.SetAuthor("John Doe")
	cp.SetCategory("Invoices")
	cp.SetContentStatus("Draft")
	cp.SetLastModifiedBy("Jane Smith")
	cp.SetCreated(time.Now())
	cp.SetModified(time.Now())
	doc.SaveToFile("document.docx")
}

func TestTemp(t *testing.T) {
	// When Word saves a document, it removes all unused styles.  This means to
	// copy the styles from an existing document, you must first create a
	// document that contains text in each style of interest.  As an example,
	// see the template.docx in this directory.  It contains a paragraph set in
	// each style that Word supports by default.
	doc, err := document.OpenTemplate("temp.docx")
	if err != nil {
		log.Fatalf("error opening Windows Word 2016 document: %s", err)
	}

	// We can now print out all styles in the document, verifying that they
	// exist.
	for _, s := range doc.Styles.Styles() {
		fmt.Println("style", s.Name(), "has ID of", s.StyleID(), "type is", s.Type())
	}

	// And create documents setting their style to the style ID (not style name).
	para := doc.AddParagraph()
	para.SetStyle("Title")
	para.AddRun().AddText("My Document Title")

	para = doc.AddParagraph()
	para.SetStyle("Subtitle")
	para.AddRun().AddText("Document Subtitle")

	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	para.AddRun().AddText("Major Section")
	para = doc.AddParagraph()
	para = doc.AddParagraph()
	for i := 0; i < 4; i++ {
		para.AddRun().AddText(lorem)
	}

	para = doc.AddParagraph()
	para.SetStyle("Heading2")
	para.AddRun().AddText("Minor Section")
	para = doc.AddParagraph()
	for i := 0; i < 4; i++ {
		para.AddRun().AddText(lorem)
	}

	// using a pre-defined table style
	table := doc.AddTable()
	table.Properties().SetWidthPercent(90)
	table.Properties().SetStyle("GridTable4-Accent1")
	look := table.Properties().TableLook()
	// these have default values in the style, so we manually turn some of them off
	look.SetFirstColumn(false)
	look.SetFirstRow(true)
	look.SetLastColumn(false)
	look.SetLastRow(true)
	look.SetHorizontalBanding(true)

	for r := 0; r < 5; r++ {
		row := table.AddRow()
		for c := 0; c < 5; c++ {
			cell := row.AddCell()
			cell.AddParagraph().AddRun().AddText(fmt.Sprintf("row %d col %d", r+1, c+1))
		}
	}
	doc.SaveToFile("use-template.docx")
}
