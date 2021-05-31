package controller

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type LineController struct {
}

func NewLineController() *LineController {
	lineCon := new(LineController)
	return lineCon
}

func (con *LineController) LineHandler(c echo.Context) (err error) {
	err = godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	bot, err := linebot.New(
		os.Getenv("LINE_SECRET_KEY"),
		os.Getenv("LINE_ACCESS_TOKEN_KEY"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	//リクエストの内容チェック
	events, err := bot.ParseRequest(c.Request())
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Response().Writer.WriteHeader(400)
		} else {
			c.Response().Writer.WriteHeader(500)
		}
		return
	}

	//イベントの内容をチェックして個別に処理していく。
	for _, event := range events {
		//eventtypeがテキスト、画像、スタンプetc...で処理を変える。
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				replyMessage := message.Text
				fmt.Println(message.Text)
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					fmt.Println(err)
				}
			case *linebot.LocationMessage:
				sendRestInfo(bot, event)
			}
		}
	}
	return
}

func sendRestInfo(bot *linebot.Client, e *linebot.Event) {
	msg := e.Message.(*linebot.LocationMessage)

	lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
	lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

	relpyMsg, err := getRestaurantInfo(lat, lng)
	if err != nil {
		fmt.Println(err)
	}

	//carouselTemplateを作成する。
	res := linebot.NewTemplateMessage("レストラン一覧画面", linebot.NewCarouselTemplate(relpyMsg...).WithImageOptions("rectangle", "cover"))

	//MessageをLineにかえす。DoでFunctionを起動。
	_, err = bot.ReplyMessage(e.ReplyToken, res).Do()
	if err != nil {
	}
}

func getRestaurantInfo(lat string, lng string) (css []*linebot.CarouselColumn, err error) {
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("http://webservice.recruit.co.jp/hotpepper/shop/v1/?key=%s&lat=%s&lng=%s&format=json", apiKey, lat, lng)

	//リクエストしてレスポンスボディーをGET
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	//Getしてきたbodyを取得する。
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var datas response
	//json形式で帰ってきたbodyをjsonUnmarshalでdata構造体に入れる。
	err = json.Unmarshal(body, &datas)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, data := range datas.Results.Shop {
		cc := linebot.NewCarouselColumn(
			data.Photo.Mobile.L,
			data.Name,
			data.Address,
			linebot.NewURIAction("ホットペッパーで開く", data.Urls.Pc),
		)
		fmt.Println(data.Photo.Mobile.L)
		css = append(css, cc)
	}
	fmt.Println(datas.Results.Shop)
	return
}

type response struct {
	Results results `json:"results"`
}

type results struct {
	Shop []shop `json:"shop"`
}

type shop struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Photo   photo  `json:"photo"`
	Urls    urls   `json:"urls"`
}

type urls struct {
	Pc string `json:"pc"`
}

type photo struct {
	Mobile mobile `json:"mobile"`
}

type mobile struct {
	L string `json:"l"`
}
