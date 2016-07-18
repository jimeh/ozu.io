// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package web

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

var _ = json.RawMessage{} // suppress unused package warning

func easyjson_6ff3ac1d_decode_github_com_jimeh_ozu_io_web_Response(in *jlexer.Lexer, out *Response) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "uid":
			out.UID = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "target":
			out.Target = string(in.String())
		case "error":
			out.Error = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjson_6ff3ac1d_encode_github_com_jimeh_ozu_io_web_Response(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"uid\":")
	out.String(string(in.UID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"url\":")
	out.String(string(in.URL))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"target\":")
	out.String(string(in.Target))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"error\":")
	out.String(string(in.Error))
	out.RawByte('}')
}
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson_6ff3ac1d_encode_github_com_jimeh_ozu_io_web_Response(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson_6ff3ac1d_encode_github_com_jimeh_ozu_io_web_Response(w, v)
}
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson_6ff3ac1d_decode_github_com_jimeh_ozu_io_web_Response(&r, v)
	return r.Error()
}
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson_6ff3ac1d_decode_github_com_jimeh_ozu_io_web_Response(l, v)
}