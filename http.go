/*
 * @Description:
 * @Author: chunhua.yang
 * @Date: 2020-05-29 12:06:33
 * @LastEditors: chunhua.yang
 * @LastEditTime: 2020-05-31 21:27:24
 */
package gonginx

import (
	"errors"
	"fmt"
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
	//parameters := directive.GetParameters()
	http := &Http{
		HttpName: "http", //first parameter of the directive is the upstream name
	}

	block := directive.GetBlock()
	if block == nil {
		return nil, errors.New("http directive must have a block")
	}

	if len(directive.GetBlock().GetDirectives()) > 0 {

		for _, d := range directive.GetBlock().GetDirectives() {
			if d.GetName() == "server" {
				s, _ := NewServer(d)
				http.Servers = append(http.Servers, s)
			}
		}
	}

	fmt.Println("==========", len(http.Servers))
	http.Block = directive.GetBlock()

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

func (h *Http) SetBlock(block IBlock) IBlock {
	h.Block = block
	return h.Block
}

func (h *Http) GetDirectives() []IDirective {
	directives := make([]IDirective, 0)

	directives = append(directives, h.Block.GetDirectives()...)

	return directives
}

func (h *Http) AddServer(server *Server) {
	//directives := make([]IDirective, 0)
	//directives = append(directives, server.Block.GetDirectives()...)
	//fmt.Println(len(directives))

	//h.Directives = append(h.Directives, directives...)
	//得到所有的节点
	directives := h.GetBlock().GetDirectives()
	//新增server块节点
	directive := &Directive{
		Name:       "server",
		Parameters: make([]string, 0),
		Block:      server.Block,
	}
	directives = append(directives, directive)

	var block = &Block{Directives: directives}
	h.SetBlock(block)
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
