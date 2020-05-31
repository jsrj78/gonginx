/*
 * @Description:
 * @Author: chunhua.yang
 * @Date: 2020-05-29 12:06:33
 * @LastEditors: chunhua.yang
 * @LastEditTime: 2020-05-31 00:38:46
 */
package gonginx

//Config  represents a whole config file.
type Config struct {
	*Block
	FilePath string
}

//TODO(tufan): move that part inti dumper package
//SaveToFile save config to a file
//TODO: add custom file / folder path support
//func (c *Config) SaveToFile(style *dumper.Style) error {
//	//wrilte file
//	dirPath := filepath.Dir(c.FilePath)
//	if _, err := os.Stat(dirPath); err != nil {
//		err := os.MkdirAll(dirPath, os.ModePerm)
//		if err != nil {
//			return err //TODO: do we reallt need to find a way to test dir creating error?
//		}
//	}
//
//	//write main file
//	err := ioutil.WriteFile(c.FilePath, c.ToByteArray(style), 0644)
//	if err != nil {
//		return err //TODO: do we need to find a way to test writing error?
//	}
//
//	//write sub files (inlude /file/path)
//	for _, directive := range c.Block.Directives {
//		if fs, ok := (interface{}(directive)).(FileDirective); ok {
//			err := fs.SaveToFile(style)
//			if err != nil {
//				return err
//			}
//		}
//	}
//
//	return nil
//}

//FindDirectives find directives from whole config block
func (c *Config) FindDirectives(directiveName string) []IDirective {
	return c.Block.FindDirectives(directiveName)
}

//FindUpstreams find directives from whole config block
func (c *Config) FindUpstreams() []*Upstream {
	var upstreams []*Upstream
	directives := c.Block.FindDirectives("upstream")
	for _, directive := range directives {
		//	up, _ := NewUpstream(directive)
		upstreams = append(upstreams, directive.(*Upstream))
	}
	return upstreams
}

func (c *Config) FindHttp() *Http {

	http := c.Block.FindDirectives("http")[0]
	// for _, directive := range directives {
	// 	//	up, _ := NewUpstream(directive)
	// 	upstreams = append(upstreams, directive.(*Upstream))
	// }
	return http.(*Http)
}
