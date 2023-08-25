package truyenfullwebapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/hthai2201/webtruyen-go/internal/entity"
)

// TruyenFullWebAPI -.
type TruyenFullWebAPI struct {
	BaseURL string
}

// New -.
func New() *TruyenFullWebAPI {

	return &TruyenFullWebAPI{
		BaseURL: "https://truyenfull.vn",
	}
}

type GetStoryChaptersRequest struct {
	Type   string `json:"type"`
	TID    int    `json:"tid"`
	TASCII string `json:"tascii"`
	TName  string `json:"tname"`
	Page   int    `json:"page"`
	TotalP int    `json:"totalp"`
}
type GetStoryChaptersResponse struct {
	ChapList string `json:"chap_list"`
	PageList string `json:"page_list"`
}

// Translate -.
func (t *TruyenFullWebAPI) GetStoryChapters(story entity.Story, opt map[string]string) (GetStoryChaptersResponse, error) {
	queryParams := url.Values{}
	queryParams.Set("type", "list_chapter")
	queryParams.Set("tid", story.TID)
	queryParams.Set("tascii", story.Slug)
	queryParams.Set("tname", story.Name)
	queryParams.Set("page", opt["page"])
	queryParams.Set("totalp", opt["totalp"])
	resp, err := http.Get("https://truyenfull.vn/ajax.php?" + queryParams.Encode())
	if err != nil {
		return GetStoryChaptersResponse{}, err
	}
	defer resp.Body.Close()
	var response GetStoryChaptersResponse
	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &response)

	return response, nil
}

func (t *TruyenFullWebAPI) GetStoryURLBySlug(slug string) string {
	return t.BaseURL + "/" + slug
}
func (t *TruyenFullWebAPI) GetChapterURLBySlug(sSlug string, cSlug string) string {
	return t.BaseURL + "/" + sSlug + "/" + cSlug
}
func (t *TruyenFullWebAPI) ExtractStorySlug(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		return ""
	}
	return parts[1]
}
