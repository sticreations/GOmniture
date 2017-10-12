package GOmniture

import "time"

type Locale string

const (
	English    Locale = "en_US"
	German     Locale = "de_DE"
	Spanish    Locale = "es_ES"
	French     Locale = "fr_FR"
	Japanese   Locale = "jp_JP"
	Portuguese Locale = "pt_BR"
	Korean     Locale = "ko_KR"
	Chinese    Locale = "zh_CN"
	Chinese_CN Locale = "zh_CN"
	Chinese_TW Locale = "zh_TW"

	en_US Locale = "en_US"
	de_DE Locale = "de_DE"
	es_ES Locale = "es_ES"
	fr_FR Locale = "fr_FR"
	jp_JP Locale = "jp_JP"
	pt_BR Locale = "pt_BR"
	ko_KR Locale = "ko_KR"
	zh_CN Locale = "zh_CN"
	zh_TW Locale = "zh_TW"
)

type SearchType string

const (
	SearchTypeAND SearchType = "AND"
	SearchTypeOR  SearchType = "OR"
	SearchTypeNOT SearchType = "NOT"
)

type ReportQuery struct {
	ReportDescription *Description `json:"reportDescription"`
}

type Description struct {
	ReportSuiteID    string     `json:"reportSuiteID"`
	Date             string     `json:"date,omitempty"`
	DateFrom         string     `json:"dateFrom,omitempty"`
	DateTo           string     `json:"dateTo,omitempty"`
	DateGranularity  string     `json:"dateGranularity,omitempty"`
	Metrics          []*Metric  `json:"metrics,omitempty"`
	Elements         []*Element `json:"elements,omitempty"`
	Locale           Locale     `json:"locale,omitempty"`
	SortBy           string     `json:"sortBy,omitempty"`
	Segments         []*Segment `json:"segments,omitempty"`
	SegmentId        string     `json:"segment_id,omitempty"`
	AnomalyDetection bool       `json:"anomalyDetection,omitempty"`
	CurrentData      bool       `json:"currentData,omitempty"`
	Expedite         bool       `json:"expedite,omitempty"`
}

type Metric struct {
	Id string `json:"id"`
}

type Element struct {
	Id             string     `json:"id"`
	Classification string     `json:"classification,omitempty"`
	Top            int        `json:"top,omitempty"`
	StartingWith   int        `json:"startingWith,omitempty"`
	Search         *Search    `json:"search,omitempty"`
	Selected       []string   `json:"selected,omitempty"`
	ParentID       string     `json:"parentID,omitempty"`
	Checkpoints    []string   `json:"checkpoints,omitempty"`
	Pattern        [][]string `json:"pattern,omitempty"`
}

type Search struct {
	Type     SearchType `json:"type"`
	Keywords []string   `json:"keywords"`
	Searches []*Search  `json:"searches,omitempty"`
}

type Segment struct {
	ID             string  `json:"id"`
	Element        string  `json:"element,omitempty"`
	Search         *Search `json:"search,omitempty"`
	Classification string  `json:"classification,omitempty"`
}

type ReportResponse struct {
	WaitSeconds   float64 `json:"waitSeconds"`
	RunSeconds    float64 `json:"runSeconds"`
	Report        *Report `json:"report"`
	TimeRetrieved time.Time
}

type Report struct {
	Type        string       `json:"type"`
	ReportSuite *ReportSuite `json:"reportSuite"`
	Period      string       `json:"period"`
	Elements    []*Element   `json:"elements"`
	Metrics     []*Metric    `json:"metrics"`
	Segments    []*Segment   `json:"segments"`
	Data        []*Data      `json:"data"`
	Totals      []float64    `json:"totals"`
}

type ReportSuite struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Data struct {
	Name           string    `json:"name"`
	Url            string    `json:"url"`
	Path           *DataPath `json:"path"`
	ParentID       string    `json:"parentID"`
	Year           int       `json:"year"`
	Month          int       `json:"month"`
	Day            int       `json:"day"`
	Hour           int       `json:"hour"`
	Counts         []float64 `json:"counts"`
	UpperBounds    []float64 `json:"upperBounds"`
	LowerBounds    []float64 `json:"lowerBounds"`
	Forecasts      []float64 `json:"forecasts"`
	BreakdownTotal []float64 `json:"breakdownTotal"`
	Breakdown      []*Data   `json:"breakdown"`
}

type DataPath struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type queueReportResponse struct {
	ReportID int `json:"reportID"`
}

type getReport struct {
	ReportID int `json:"reportID"`
}
