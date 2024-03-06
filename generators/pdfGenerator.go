package generators

import (
	"strconv"
	"time"

	"github.com/andreabertanzon/simtek-generate/models"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	marotoCore "github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type PdfGenerator struct {
	Style Style
}

func NewPdfGenerator(*Style) Generator {
	return &PdfGenerator{
		Style: Style{
			Font:            "Arial",
			Fill:            "solid",
			TextColor:       "black",
			FontH1Dimension: 20,
			FontH2Dimension: 18,
			FontH3Dimension: 16,
			PrimaryColor:    "#000000",
			SecondaryColor:  "#FFC0CB",
		},
	}
}

func (*PdfGenerator) Generate(interventions []models.Intervention) error {
	timestamp := time.Now().Format("02-01-2006")
	filename := "interventions-" + timestamp + ".pdf"

	m := *getMaroto(interventions)
	document, err := m.Generate()
	if err != nil {
		return err
	}

	return document.Save(filename)
}

// you can setup your default page here
func getMaroto(interventions []models.Intervention) *marotoCore.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber("{current} / {total}", props.Bottom).
		Build()

	m := maroto.New(cfg)
	// deco := maroto.NewMetricsDecorator(m)

	// colStyle := &props.Cell{
	// 	BackgroundColor: &props.Color{Red: 80, Green: 80, Blue: 80},
	// 	BorderType:      border.Full,
	// 	BorderColor:     &props.Color{Red: 200, Green: 0, Blue: 0},
	// 	LineStyle:       linestyle.Dashed,
	// 	BorderThickness: 0.5,
	// }

	rowStyles := []*props.Cell{
		{
			BackgroundColor: &props.Color{Red: 220, Green: 220, Blue: 220},
			BorderType:      border.None,
			BorderColor:     &props.Color{Red: 0, Green: 0, Blue: 200},
		},
		{
			BackgroundColor: &props.Color{Red: 220, Green: 220, Blue: 220},
			BorderType:      border.Full,
			BorderColor:     &props.Color{Red: 0, Green: 0, Blue: 200},
		},
		{
			BackgroundColor: &props.Color{Red: 220, Green: 220, Blue: 220},
			BorderType:      border.Left,
			BorderColor:     &props.Color{Red: 0, Green: 0, Blue: 200},
		},
		{
			BackgroundColor: &props.Color{Red: 220, Green: 220, Blue: 220},
			BorderType:      border.Right,
			BorderColor:     &props.Color{Red: 0, Green: 0, Blue: 200},
		},
		{
			BackgroundColor: &props.Color{Red: 220, Green: 220, Blue: 220},
			BorderType:      border.Top,
			BorderColor:     &props.Color{Red: 0, Green: 0, Blue: 200},
		},
		{
			BackgroundColor: &props.Color{Red: 220, Green: 220, Blue: 220},
			BorderType:      border.Bottom,
			BorderColor:     &props.Color{Red: 0, Green: 0, Blue: 200},
		},
	}

	m.AddRow(10,
		text.NewCol(4, "Data: 26-02-1989", props.Text{
			Top:  2,
			Left: 2,
			Size: 14,
		}).WithStyle(rowStyles[1]),
		text.NewCol(8, "Totale interventi: 10", props.Text{
			Top:  2,
			Left: 2,
			Size: 14,
		}).WithStyle(rowStyles[1]),
	)

	for _, intervention := range interventions {
		m.AddRow(10, text.NewCol(12, intervention.Site, props.Text{
			Size: 18,
		}))
		m.AddRow(10, text.NewCol(12, "Operatori:"))
		for key, value := range models.Flatten(intervention.Workers) {
			m.AddRow(10, text.NewCol(8, key), text.NewCol(4, strconv.Itoa(value)))
		}
		m.AddRow(10,
			line.NewCol(2, props.Line{
				Thickness: 2,
			}),
		)

	}

	return &m
}
