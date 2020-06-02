package httpsource

import (
    "fmt"
    "time"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strconv"

    "github.com/HoMuChen/go-park/domain"
)

type service struct {
    baseUrl     string
}

func New() domain.PriceHttpSource {
    return &service{
        "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=%s&json=1",
    }
}

func (s *service) FetchRealtime(company domain.Company) (domain.Price, error) {
    code := company.Code
    url := s.genUrl(code)

    res, err := http.Get(url)
    if err != nil {
        return domain.Price{}, err
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return domain.Price{}, err
    }

    var result struct {
        MsgArray    []struct {
            ID      string `json:"c"`
            Name    string `json:"n"`
            Price   string `json:"z"`
            Time    string `json:"tlong"`
        }
    }
    err = json.Unmarshal(body, &result)
    if err != nil {
        return domain.Price{}, err
    }

    if len(result.MsgArray) == 0 {
        return domain.Price{}, domain.ErrBadParamInput
    }

    datetime, err := strconv.ParseInt(result.MsgArray[0].Time, 10, 64)
    if err != nil {
        return domain.Price{}, err
    }
    if result.MsgArray[0].Price == "-" {
        return domain.Price{company, time.Unix(datetime/1000, 0), -1}, nil
    }
    value, err := strconv.ParseFloat(result.MsgArray[0].Price, 64)
    if err != nil {
        return domain.Price{}, err
    }
    return domain.Price{company, time.Unix(datetime/1000, 0), value}, nil
}

func (s *service) genUrl(id string) string {
    return fmt.Sprintf(s.baseUrl, id)
}
