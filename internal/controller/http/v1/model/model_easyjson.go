// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC80ae7adDecodeGithubComPprishchepaGoInvitecoderExampleInternalControllerHttpV1Model(in *jlexer.Lexer, out *AcceptInviteRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "email":
			out.Email = string(in.String())
		case "code":
			out.Code = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComPprishchepaGoInvitecoderExampleInternalControllerHttpV1Model(out *jwriter.Writer, in AcceptInviteRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		out.String(string(in.Code))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AcceptInviteRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComPprishchepaGoInvitecoderExampleInternalControllerHttpV1Model(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AcceptInviteRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComPprishchepaGoInvitecoderExampleInternalControllerHttpV1Model(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AcceptInviteRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComPprishchepaGoInvitecoderExampleInternalControllerHttpV1Model(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AcceptInviteRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComPprishchepaGoInvitecoderExampleInternalControllerHttpV1Model(l, v)
}
