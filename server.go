/*
 * @Description:
 * @Author: chunhua.yang
 * @Date: 2020-05-29 12:06:33
 * @LastEditors: chunhua.yang
 * @LastEditTime: 2020-05-31 15:42:39
 */
package gonginx

import (
	"errors"
)

//Server represents server block
type Server struct {
	Block      IBlock
	Directives []IDirective
}

//NewServer create a server block from a directive with block
func NewServer(directive IDirective) (*Server, error) {
	if block := directive.GetBlock(); block != nil {
		return &Server{
			Block: block,
		}, nil
	}

	return nil, errors.New("server directive must have a block")
}

//GetName get directive name to construct the statment string
func (s *Server) GetName() string { //the directive name.
	return "server"
}

//GetParameters get directive parameters if any
func (s *Server) GetParameters() []string {
	return []string{}
}

//GetBlock get block if any
func (s *Server) GetBlock() IBlock {
	return s.Block
}

func (s *Server) GetDirective() *Directive {
	//First, generate a new directive from upstream server
	directive := &Directive{
		Name:       "server",
		Parameters: make([]string, 0),
		Block:      s.Block,
	}

	return directive
}

func (s *Server) GetDirectives() []IDirective {
	directives := make([]IDirective, 0)
	directives = append(directives, s.Directives...)
	// for _, ss := range s.Servers {
	// 	directives = append(directives, ss)
	// }

	return directives
}

func (s *Server) FindDirectives(directiveName string) []IDirective {
	directives := make([]IDirective, 0)
	for _, directive := range s.Directives {
		if directive.GetName() == directiveName {
			directives = append(directives, directive)
		}
		if directive.GetBlock() != nil {
			directives = append(directives, directive.GetBlock().FindDirectives(directiveName)...)
		}
	}

	return directives
}
