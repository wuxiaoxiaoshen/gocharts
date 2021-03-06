package charts

import (
	"encoding/json"
	"net/http"

	"github.com/wuxiaoxiaoshen/gocharts/charts/options"
)

type Charts interface {
	Plot() func(writer http.ResponseWriter, request *http.Request)
	Save(fileName string)
	TypeName() string
}

const (
	defaultAlpha       = 0.9
	defaultBorderWidth = 1
)

type Base struct {
	Type    string   `json:"type"`
	Title   []string `json:"title"`
	Data    `json:"data"`
	Options `json:"options"`
}

type Data struct {
	Labels   []interface{} `json:"labels"`
	DataSets []DataSets    `json:"datasets"`
}

type DataSets struct {
	Type            string        `json:"type,omitempty"`
	Label           string        `json:"label,omitempty"`
	Data            []interface{} `json:"data"`
	BackgroundColor []string      `json:"backgroundColor,omitempty"`
	BorderColor     []string      `json:"borderColor,omitempty"`
	BorderWidth     int           `json:"borderWidth",default:"1"`
	Fill            bool          `json:"fill"`
	defaultAlpha    float64
}

type Options map[string]interface{}

func (base *Base) SetLabels(labels []interface{}) {
	base.Data.Labels = labels
}

func (base *Base) AddDataSet(dataSets ...DataSets) {
	base.DataSets = append(base.DataSets, dataSets...)
}

func (base *Base) AddOptions(key string, values interface{}) {
	base.Options[key] = values
}

func (base Base) NewDataSet(label string, data []interface{}) *DataSets {
	color, borderColor := options.RandomColorBoth(len(data), defaultAlpha)

	dataSets := &DataSets{
		Label:           label,
		Data:            data,
		BackgroundColor: color,
		BorderColor:     borderColor,
		BorderWidth:     defaultBorderWidth,
	}
	return dataSets
}

func (base *Base) JsonMarshal() ([]byte, error) {
	return json.Marshal(base)
}

func (base Base) ChartsOpts() ([]byte, error) {
	return base.JsonMarshal()
}

func (base *Base) setColor() {
	for index, i := range base.DataSets {
		i.defaultAlpha = defaultAlpha
		number := len(i.Data)
		base.DataSets[index].BorderWidth = defaultBorderWidth
		base.DataSets[index].BackgroundColor, base.DataSets[index].BorderColor = options.RandomColorBoth(number, i.defaultAlpha)
	}
}

func (base *Base) setBorderWidth(width int) {
	var borderWidth int
	if width <= 0 {
		borderWidth = defaultBorderWidth
	} else {
		borderWidth = width
	}
	for index := range base.DataSets {
		base.DataSets[index].BorderWidth = borderWidth
	}
}

func (base *Base) TypeName() string {
	return base.Type
}

func (base *Base) Save(fileName string) {}
