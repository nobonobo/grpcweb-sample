package main

import (
	"context"
	"log"
	"net/url"
	"strings"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/html/htmlevent"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	web "github.com/nobonobo/grpcweb-sample/proto"
)

// ParseHash ...
func ParseHash(s string) *url.URL {
	orig, _ := url.Parse(s)
	u, _ := url.Parse(orig.Fragment)
	if u.Path == "" {
		u.Path = "/"
	}
	return u
}

// GetURL ...
func GetURL() *url.URL {
	return ParseHash(webapi.GetWindow().Location().Hash())
}

// TopView ...
type TopView struct {
	vecty.Core
	pages map[string]vecty.Component
}

// AddPage ...
func (c *TopView) AddPage(prefix string, p vecty.Component) {
	if c.pages == nil {
		c.pages = map[string]vecty.Component{}
	}
	c.pages[prefix] = p
}

// GetPage ...
func (c *TopView) GetPage(s string) vecty.Component {
	if c.pages == nil {
		return nil
	}
	bestPrefix := ""
	var bestComponent vecty.Component
	for prefix, component := range c.pages {
		if strings.HasPrefix(s, prefix) && len(s) > len(bestPrefix) {
			bestPrefix = prefix
			bestComponent = component
		}
	}
	return bestComponent
}

func (c *TopView) onHashChange(event *domcore.Event) {
	ev := htmlevent.HashChangeEventFromJS(event.JSValue())
	log.Println(ev.OldURL(), "->", ev.NewURL())
	vecty.Rerender(c)
}

// Render ...
func (c *TopView) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("TopView")
	return elem.Body(
		elem.Header(
			vecty.Markup(
				vecty.Class("navbar"),
				vecty.Style("box-shadow", "2px 2px 2px lightgrey"),
				vecty.Style("padding", "1rem"),
			),
			elem.Section(
				vecty.Markup(
					vecty.Class("navbar-section"),
				),
				elem.Anchor(
					vecty.Markup(
						prop.Href("#/"),
						vecty.Class("navbar-brand"),
					),
					elem.Span(
						vecty.Markup(
							vecty.Class("text-large"),
						),
						vecty.Text("Sample1"),
					),
				),
				elem.Anchor(
					vecty.Markup(
						prop.Href("#/next"),
						vecty.Class("btn", "btn-link"),
					),
					vecty.Text("Next"),
				),
			),
			elem.Section(
				vecty.Markup(
					vecty.Class("navbar-section"),
				),
				elem.Anchor(
					vecty.Markup(
						prop.Href("#/about"),
						vecty.Class("btn", "btn-link"),
					),
					vecty.Text("About"),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("container"),
				vecty.Style("padding", "1rem"),
			),
			c.GetPage(GetURL().Path),
		),
	)
}

// Home ...
type Home struct {
	vecty.Core
}

// Render ...
func (c *Home) Render() vecty.ComponentOrHTML {
	return elem.Heading3(
		vecty.Text("Hello!"),
	)
}

// Next ...
type Next struct {
	vecty.Core
}

// Render ...
func (c *Next) Render() vecty.ComponentOrHTML {
	return elem.Heading3(
		vecty.Text("World!"),
	)
}

// About ...
type About struct {
	vecty.Core
}

// Render ...
func (c *About) Render() vecty.ComponentOrHTML {
	return elem.Heading3(
		vecty.Text("About"),
	)
}

func main() {
	meta := webapi.GetDocument().CreateElement("meta", nil)
	meta.SetAttribute("name", "viewport")
	meta.SetAttribute("content", "width=device-width,initial-scale=1")
	webapi.GetDocument().Head().Append(dom.UnionFromJS(meta.JSValue()))
	top := &TopView{}
	webapi.GetWindow().AddEventListener("hashchange",
		domcore.NewEventListenerFunc(top.onHashChange), nil)
	cc, err := grpc.Dial("")
	if err != nil {
		grpclog.Println(err)
		return
	}
	client := web.NewBackendClient(cc)
	resp, err := client.GetUser(context.Background(), &web.GetUserRequest{
		UserId: "1234",
	})
	log.Println(resp)
	top.AddPage("/", &Home{})
	top.AddPage("/next", &Next{})
	top.AddPage("/about", &About{})
	vecty.AddStylesheet("https://unpkg.com/spectre.css/dist/spectre.min.css")
	vecty.AddStylesheet("https://unpkg.com/spectre.css/dist/spectre-exp.min.css")
	vecty.AddStylesheet("https://unpkg.com/spectre.css/dist/spectre-icons.min.css")
	vecty.RenderBody(top)
}
