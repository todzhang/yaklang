package schema

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/yaklang/yaklang/common/utils"
)

const COLORPREFIX = "YAKIT_COLOR_"

func yakitColor(i string) string {
	return COLORPREFIX + i
}

type HTTPFlow struct {
	gorm.Model

	HiddenIndex        string `gorm:"index"`
	NoFixContentLength bool   `json:"no_fix_content_length"`
	Hash               string `gorm:"unique_index"`
	IsHTTPS            bool
	Url                string `gorm:"index"`
	Path               string
	Method             string
	BodyLength         int64
	ContentType        string
	StatusCode         int64
	SourceType         string
	Request            string
	Response           string
	Duration           int64
	GetParamsTotal     int
	PostParamsTotal    int
	CookieParamsTotal  int
	IPAddress          string
	RemoteAddr         string
	IPInteger          int
	Tags               string // 用来打标！
	Payload            string

	// Websocket 相关字段
	IsWebsocket bool
	// 用来计算 websocket hash, 每次连接都不一样，一般来说，内部对象 req 指针足够了
	WebsocketHash string

	RuntimeId   string
	FromPlugin  string
	ProcessName sql.NullString

	// friendly for gorm build instance, not for store
	// 这两个字段不参与数据库存储，但是在序列化的时候，会被覆盖
	// 主要用来标记用户的 Request 和 Response 是否超大
	IsRequestOversize  bool `gorm:"-"`
	IsResponseOversize bool `gorm:"-"`

	IsTooLargeResponse         bool
	TooLargeResponseHeaderFile string
	TooLargeResponseBodyFile   string
	// 同步到企业端
	UploadOnline bool `json:"upload_online"`
}

func (f *HTTPFlow) GetRequest() string {
	unquoted, err := strconv.Unquote(f.Request)
	if err != nil {
		return ""
	}
	return unquoted
}

func (f *HTTPFlow) GetResponse() string {
	unquoted, err := strconv.Unquote(f.Response)
	if err != nil {
		return ""
	}
	return unquoted
}

func (f *HTTPFlow) SetRequest(req string) {
	f.Request = strconv.Quote(req)
}

func (f *HTTPFlow) SetResponse(rsp string) {
	f.Response = strconv.Quote(rsp)
}

// 颜色与 Tag API
func (f *HTTPFlow) AddTag(appendTags ...string) {
	existed := utils.PrettifyListFromStringSplited(f.Tags, "|")
	existedCount := len(existed)
	extLen := len(appendTags)
	tags := make([]string, existedCount+extLen)
	copy(tags, existed)
	for i := 0; i < extLen; i++ {
		tags[i+existedCount] = appendTags[i]
	}
	f.Tags = strings.Join(utils.RemoveRepeatStringSlice(tags), "|")
}

func (f *HTTPFlow) AddTagToFirst(appendTags ...string) {
	existed := utils.PrettifyListFromStringSplited(f.Tags, "|")
	f.Tags = strings.Join(utils.RemoveRepeatStringSlice(append(appendTags, existed...)), "|")
}

func (f *HTTPFlow) HasColor(color string) bool {
	return utils.StringArrayContains(utils.PrettifyListFromStringSplited(f.Tags, "|"), color)
}

var (
	FLOW_COLOR_RED    = yakitColor("RED")
	FLOW_COLOR_GREEN  = yakitColor("GREEN")
	FLOW_COLOR_BLUE   = yakitColor("BLUE")
	FLOW_COLOR_YELLOW = yakitColor("YELLOW")
	FLOW_COLOR_ORANGE = yakitColor("ORANGE")
	FLOW_COLOR_PURPLE = yakitColor("PURPLE")
	FLOW_COLOR_CYAN   = yakitColor("CYAN")
	FLOW_COLOR_GREY   = yakitColor("GREY")
)

func (f *HTTPFlow) Red() {
	if f.HasColor(FLOW_COLOR_RED) {
		return
	}
	f.AddTag(FLOW_COLOR_RED)
}

func (f *HTTPFlow) Green() {
	if f.HasColor(FLOW_COLOR_GREEN) {
		return
	}
	f.AddTag(FLOW_COLOR_GREEN)
}

func (f *HTTPFlow) Blue() {
	if f.HasColor(FLOW_COLOR_BLUE) {
		return
	}
	f.AddTag(FLOW_COLOR_BLUE)
}

func (f *HTTPFlow) Yellow() {
	if f.HasColor(FLOW_COLOR_YELLOW) {
		return
	}
	f.AddTag(FLOW_COLOR_YELLOW)
}

func (f *HTTPFlow) Orange() {
	if f.HasColor(FLOW_COLOR_ORANGE) {
		return
	}
	f.AddTag(FLOW_COLOR_ORANGE)
}

func (f *HTTPFlow) Purple() {
	if f.HasColor(FLOW_COLOR_PURPLE) {
		return
	}
	f.AddTag(FLOW_COLOR_PURPLE)
}

func (f *HTTPFlow) Cyan() {
	if f.HasColor(FLOW_COLOR_CYAN) {
		return
	}
	f.AddTag(FLOW_COLOR_CYAN)
}

func (f *HTTPFlow) Grey() {
	if f.HasColor(FLOW_COLOR_GREY) {
		return
	}
	f.AddTag(FLOW_COLOR_GREY)
}

func (f *HTTPFlow) ColorSharp(rgbHex string) {
	if f.HasColor(yakitColor(rgbHex)) {
		return
	}
	f.AddTag(yakitColor(rgbHex))
}

func (f *HTTPFlow) CalcHash() string {
	return utils.CalcSha1(f.IsHTTPS, f.Url, f.Path, f.Method, f.BodyLength, f.ContentType, f.StatusCode, f.SourceType, f.Tags, f.Request, f.HiddenIndex, f.RuntimeId, f.FromPlugin)
}

func (f *HTTPFlow) CalcCacheHash(full bool) string {
	return utils.CalcSha1(f.ID, f.Hash, full)
}

func (f *HTTPFlow) BeforeSave() error {
	f.fixURL()
	f.Hash = f.CalcHash()
	return nil
}

func (f *HTTPFlow) fixURL() {
	urlIns := utils.ParseStringToUrl(f.Url)
	if f.IsHTTPS {
		urlIns.Scheme = "https"
	}
	if urlIns != nil {
		host, port, _ := utils.ParseStringToHostPort(urlIns.Host)
		if (port == 443 && urlIns.Scheme == "https") || (port == 80 && urlIns.Scheme == "http") {
			urlIns.Host = host
			f.Url = urlIns.String()
		}
	}
}

func (f *HTTPFlow) AfterCreate(tx *gorm.DB) (err error) {
	broadcastData.Call("httpflow", "create")
	return nil
}

func (f *HTTPFlow) AfterUpdate(tx *gorm.DB) (err error) {
	broadcastData.Call("httpflow", "update")
	return nil
}

func (f *HTTPFlow) AfterDelete(tx *gorm.DB) (err error) {
	broadcastData.Call("httpflow", "delete")
	return nil
}
