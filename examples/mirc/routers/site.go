package routers

import "github.com/alimy/mir/v2"

// Site mir's struct tag define
type Site struct {
	Chain    mir.Chain `mir:"-"`
	Group    mir.Group `mir:"v1"`
	Index    mir.Get   `mir:"/index/"`
	Articles mir.Get   `mir:"/articles/:category/"`
}
