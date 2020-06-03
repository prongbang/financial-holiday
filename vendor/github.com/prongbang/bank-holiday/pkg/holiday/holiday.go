package holiday

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const url = "https://www.bot.or.th/Thai/FinancialInstitutions/FIholiday/Pages/HolidayCalendar.aspx?y=%s"
const caleClass = ".cal_month"
const monthID = "#ctl00_ctl73_g_2d5ecfc7_5e54_499d_8285_2fff313ce83f_ctl00_MonthLabel%d"

type Holidays map[string]map[string]string

type Utility interface {
	GetFinancialHoliday(year string) Holidays
}

type utility struct {
}

func (u *utility) GetFinancialHoliday(year string) Holidays {
	holidayURL := fmt.Sprintf(url, year)

	// Request the HTML page.
	res, err := http.Get(holidayURL)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	calendarDom := doc.Find(caleClass)

	data := Holidays{}
	calendarDom.Find(".row").Each(func(i int, s *goquery.Selection) {
		s.Find(".col-xs-12 .col-sm-12 .col-md-4").Each(func(i int, col *goquery.Selection) {
			month := map[string]string{}
			title := col.Find(".cal_month_text").Text()
			col.Find("table > tbody > tr > td").Each(func(i int, td *goquery.Selection) {
				dayTitle := td.AttrOr("title", "")
				if dayTitle != "" {
					day := td.Text()
					month[day] = dayTitle
				}
			})
			data[monthMapping(title)] = month
		})
	})
	return data
}

func monthMapping(month string) string {
	switch month {
	case "มกราคม":
		return "january"
	case "กุมภาพันธ์":
		return "february"
	case "มีนาคม":
		return "march"
	case "เมษายน":
		return "april"
	case "พฤษภาคม":
		return "may"
	case "มิถุนายน":
		return "june"
	case "กรกฎาคม":
		return "july"
	case "สิงหาคม":
		return "august"
	case "กันยายน":
		return "september"
	case "ตุลาคม":
		return "october"
	case "พฤศจิกายน":
		return "november"
	case "ธันวาคม":
		return "december"
	default:
		return ""
	}
}

func NewUtility() Utility {
	return &utility{}
}
