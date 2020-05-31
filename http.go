/*
 * @Description:
 * @Author: chunhua.yang
 * @Date: 2020-05-29 12:06:33
 * @LastEditors: chunhua.yang
 * @LastEditTime: 2020-05-31 00:25:50
 */
package gonginx

import (
	"errors"
)

//Http represents http block
type Http struct {
	Block      IBlock
	HttpName   string
	Servers    []*Server
	Directives []IDirective
}

//NewHttp create an http block from a directive which has a block
func NewHttp(directive IDirective) (*Http, error) {
	parameters := directive.GetParameters()
	http := &Http{
		HttpName: parameters[0], //first parameter of the directive is the upstream name
	}

	block := directive.GetBlock()
	if block == nil {
		return nil, errors.New("http directive must have a block")
	}

	if len(directive.GetBlock().GetDirectives()) > 0 {
		for _, d := range directive.GetBlock().GetDirectives() {
			if d.GetName() == "server" {
				s, _ := NewServer(d)
				http.Servers = append(http.Servers,s)
			}
		}
	}

	// return &Http{
	// 	Block: block,
	// }, nil
	return http, nil

}

//GetName get directive name to construct the statment string
func (h *Http) GetName() string { //the directive name.
	return "http"
}

//GetParameters get directive parameters if any
func (h *Http) GetParameters() []string {
	return []string{}
}

//GetBlock get block if any
func (h *Http) GetBlock() IBlock {
	return h.Block
}

func (h *Http) GetDirectives() []IDirective {
	directives := make([]IDirective, 0)
	directives = append(directives, h.Directives...)
	for _, ss := range h.Servers {
		directives = append(directives, ss)
	}

	return directives
}

func (h *Http) AddServer(server *Server) {
	h.Servers = append(h.Servers, server)
}

func (h *Http) FindDirectives(directiveName string) []IDirective {
	directives := make([]IDirective, 0)
	for _, directive := range h.Directives {
		if directive.GetName() == directiveName {
			directives = append(directives, directive)
		}
		if directive.GetBlock() != nil {
			directives = append(directives, directive.GetBlock().FindDirectives(directiveName)...)
		}
	}

	return directives
}
