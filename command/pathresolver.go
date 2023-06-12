package command

func Resolve(r *Request) string {

	profile := r.Profile
	params := r.Params

	return "http://" + profile.Host + ":" + profile.Port + "/" + params.Command + "?app=" + profile.App
}
