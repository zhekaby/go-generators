// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package tests

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

func easyjsonD2b7633eDecodeGithubComZhekabyGoGeneratorMongoRequestwraperTests(in *jlexer.Lexer, out *device) {
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
		case "user_id":
			out.UserID = string(in.String())
		case "locale":
			out.Locale = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "native_push_token":
			out.NativePushToken = string(in.String())
		case "native_voip_token":
			out.NativeVoIPToken = string(in.String())
		case "carrier":
			out.Carrier = string(in.String())
		case "mcc":
			out.Mcc = string(in.String())
		case "mnc":
			out.Mnc = string(in.String())
		case "os_version":
			out.OsVersion = string(in.String())
		case "build_number":
			out.BuildNumber = string(in.String())
		case "app_version":
			out.AppVersion = string(in.String())
		case "country_code":
			out.CountryCode = string(in.String())
		case "phone_number":
			out.PhoneNumber = string(in.String())
		case "mode":
			out.Mode = string(in.String())
		case "flags":
			if in.IsNull() {
				in.Skip()
				out.Flags = nil
			} else {
				in.Delim('[')
				if out.Flags == nil {
					if !in.IsDelim(']') {
						out.Flags = make([]*flag, 0, 8)
					} else {
						out.Flags = []*flag{}
					}
				} else {
					out.Flags = (out.Flags)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *flag
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(flag)
						}
						easyjsonD2b7633eDecodeGithubComZhekabyGoGeneratorMongoRequestwraperTests1(in, v1)
					}
					out.Flags = append(out.Flags, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "MyData":
			if in.IsNull() {
				in.Skip()
				out.MyData = nil
			} else {
				if out.MyData == nil {
					out.MyData = new(data)
				}
				easyjsonD2b7633eDecodeGithubComZhekabyGoGeneratorMongoRequestwraperTests2(in, out.MyData)
			}
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
func easyjsonD2b7633eEncodeGithubComZhekabyGoGeneratorMongoRequestwraperTests(out *jwriter.Writer, in device) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserID))
	}
	if in.Locale != "" {
		const prefix string = ",\"locale\":"
		out.RawString(prefix)
		out.String(string(in.Locale))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"native_push_token\":"
		out.RawString(prefix)
		out.String(string(in.NativePushToken))
	}
	{
		const prefix string = ",\"native_voip_token\":"
		out.RawString(prefix)
		out.String(string(in.NativeVoIPToken))
	}
	if in.Carrier != "" {
		const prefix string = ",\"carrier\":"
		out.RawString(prefix)
		out.String(string(in.Carrier))
	}
	if in.Mcc != "" {
		const prefix string = ",\"mcc\":"
		out.RawString(prefix)
		out.String(string(in.Mcc))
	}
	if in.Mnc != "" {
		const prefix string = ",\"mnc\":"
		out.RawString(prefix)
		out.String(string(in.Mnc))
	}
	if in.OsVersion != "" {
		const prefix string = ",\"os_version\":"
		out.RawString(prefix)
		out.String(string(in.OsVersion))
	}
	if in.BuildNumber != "" {
		const prefix string = ",\"build_number\":"
		out.RawString(prefix)
		out.String(string(in.BuildNumber))
	}
	if in.AppVersion != "" {
		const prefix string = ",\"app_version\":"
		out.RawString(prefix)
		out.String(string(in.AppVersion))
	}
	if in.CountryCode != "" {
		const prefix string = ",\"country_code\":"
		out.RawString(prefix)
		out.String(string(in.CountryCode))
	}
	if in.PhoneNumber != "" {
		const prefix string = ",\"phone_number\":"
		out.RawString(prefix)
		out.String(string(in.PhoneNumber))
	}
	if in.Mode != "" {
		const prefix string = ",\"mode\":"
		out.RawString(prefix)
		out.String(string(in.Mode))
	}
	{
		const prefix string = ",\"flags\":"
		out.RawString(prefix)
		if in.Flags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Flags {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjsonD2b7633eEncodeGithubComZhekabyGoGeneratorMongoRequestwraperTests1(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"MyData\":"
		out.RawString(prefix)
		if in.MyData == nil {
			out.RawString("null")
		} else {
			easyjsonD2b7633eEncodeGithubComZhekabyGoGeneratorMongoRequestwraperTests2(out, *in.MyData)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v device) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComZhekabyGoGeneratorMongoRequestwraperTests(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v device) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComZhekabyGoGeneratorMongoRequestwraperTests(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *device) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComZhekabyGoGeneratorMongoRequestwraperTests(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *device) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComZhekabyGoGeneratorMongoRequestwraperTests(l, v)
}
func easyjsonD2b7633eDecodeGithubComZhekabyGoGeneratorMongoRequestwraperTests2(in *jlexer.Lexer, out *data) {
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
func easyjsonD2b7633eEncodeGithubComZhekabyGoGeneratorMongoRequestwraperTests2(out *jwriter.Writer, in data) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}
func easyjsonD2b7633eDecodeGithubComZhekabyGoGeneratorMongoRequestwraperTests1(in *jlexer.Lexer, out *flag) {
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
		case "type":
			out.Type = string(in.String())
		case "event":
			out.Event = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComZhekabyGoGeneratorMongoRequestwraperTests1(out *jwriter.Writer, in flag) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"event\":"
		out.RawString(prefix)
		out.String(string(in.Event))
	}
	out.RawByte('}')
}