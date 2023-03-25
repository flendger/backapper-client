package pathresolver

import "backapper-client/params"

func Resolve(p *params.Params) string {
	return "http://" + p.Address + "/" + p.Command + "?app=" + p.AppName
}
