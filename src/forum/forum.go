package forum

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf/browser"
	"gopkg.in/headzoo/surf.v1"
)

const forumURL = "http://slitherin.potterforum.ru"
const loginHandler = "login.php"
const getTopic = "viewtopic.php"

//Forum struct
type Forum struct {
	browser    *browser.Browser
	TopicTitle string
	postsRaw   []*goquery.Selection
	parsed     []string
}

//Init browser
func (f *Forum) Init() *Forum {
	f.browser = surf.NewBrowser()
	f.postsRaw = make([]*goquery.Selection, 1)
	f.browser.Open("http://slitherin.potterforum.ru/login.php?action=in")
	return f
}

//TopicText returning text from all posts of topic as one string
func (f *Forum) TopicText() string {
	return strings.Join(f.parsed, "\n")
}

//Auth cookies
func (f *Forum) Auth(login string, passwd string) (cookies []*http.Cookie) {
	cookies = f.prepareAuthData(login, passwd)
	return cookies
}

//GetTopic - method for request to fanfic page
func (f *Forum) GetTopic(topicID int) {
	// var postList = make(map[int]string)
	f.getPage(f.prepareTopicRequest(topicID))
	f.parsePosts()
}

//Prepare POST request for auth
func (f *Forum) prepareAuthData(login string, passwd string) []*http.Cookie {
	form, err := f.browser.Form("form#login")
	if err != nil {
		panic(err)
	}
	form.Input("req_username", login)
	form.Input("req_password", passwd)

	formerr := form.Submit()
	if formerr != nil {
		panic(formerr)
	}
	return f.browser.SiteCookies()
}

//Prepare GET request for topic
func (f *Forum) prepareTopicRequest(topicID int) (url string) {
	url = fmt.Sprintf("%s/%s?id=%d", forumURL, getTopic, topicID)
	return
}

func (f *Forum) getPage(url string) {
	f.browser.Open(url)
	if f.TopicTitle == "" {
		f.TopicTitle = strings.Split(ToUTF(f.browser.Title()), "~")[0]
	}
	pageLinks := f.browser.Find("a.next")
	nextPage := ""
	pageLinks.Each(func(_ int, page *goquery.Selection) {
		nextPage, _ = page.Attr("href")
	})
	f.postsRaw = append(f.postsRaw, f.browser.Find("div.post-content"))

	if nextPage != "" {
		f.getPage(nextPage)
	}
}

func (f *Forum) parsePosts() {
	for i := 0; i < len(f.postsRaw); i++ {
		var postForPage *goquery.Selection = f.postsRaw[i]
		if postForPage != nil {
			postForPage.Each(func(index int, s *goquery.Selection) {
				rawForumPost, _ := s.Html()
				// rawForumPost := s.Text()
				rawForumPost = strings.Replace(rawForumPost, "     ", " ", -1)
				decoded := ToUTF(rawForumPost)
				// decoded := rawForumPost
				f.parsed = append(f.parsed, decoded)
			})
		}
	}
}

//GetParsedPosts get
func (f *Forum) GetParsedPosts() []string {
	return f.parsed
}
