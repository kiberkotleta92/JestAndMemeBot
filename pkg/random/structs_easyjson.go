// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package random

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

func easyjson6a975c40DecodeJestRandomOrg(in *jlexer.Lexer, out *RandomResponse) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int64(in.Int64())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		case "result":
			easyjson6a975c40Decode(in, &out.Result)
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
func easyjson6a975c40EncodeJestRandomOrg(out *jwriter.Writer, in RandomResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		out.RawString(prefix)
		out.String(string(in.Jsonrpc))
	}
	{
		const prefix string = ",\"result\":"
		out.RawString(prefix)
		easyjson6a975c40Encode(out, in.Result)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RandomResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeJestRandomOrg(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RandomResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeJestRandomOrg(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RandomResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeJestRandomOrg(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RandomResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeJestRandomOrg(l, v)
}
func easyjson6a975c40Decode(in *jlexer.Lexer, out *struct {
	AdvisoryDelay int64 `json:"advisoryDelay"`
	BitsLeft      int64 `json:"bitsLeft"`
	BitsUsed      int64 `json:"bitsUsed"`
	Random        struct {
		CompletionTime string  `json:"completionTime"`
		Data           []int64 `json:"data"`
	} `json:"random"`
	RequestsLeft int64 `json:"requestsLeft"`
}) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "advisoryDelay":
			out.AdvisoryDelay = int64(in.Int64())
		case "bitsLeft":
			out.BitsLeft = int64(in.Int64())
		case "bitsUsed":
			out.BitsUsed = int64(in.Int64())
		case "random":
			easyjson6a975c40Decode1(in, &out.Random)
		case "requestsLeft":
			out.RequestsLeft = int64(in.Int64())
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
func easyjson6a975c40Encode(out *jwriter.Writer, in struct {
	AdvisoryDelay int64 `json:"advisoryDelay"`
	BitsLeft      int64 `json:"bitsLeft"`
	BitsUsed      int64 `json:"bitsUsed"`
	Random        struct {
		CompletionTime string  `json:"completionTime"`
		Data           []int64 `json:"data"`
	} `json:"random"`
	RequestsLeft int64 `json:"requestsLeft"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"advisoryDelay\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.AdvisoryDelay))
	}
	{
		const prefix string = ",\"bitsLeft\":"
		out.RawString(prefix)
		out.Int64(int64(in.BitsLeft))
	}
	{
		const prefix string = ",\"bitsUsed\":"
		out.RawString(prefix)
		out.Int64(int64(in.BitsUsed))
	}
	{
		const prefix string = ",\"random\":"
		out.RawString(prefix)
		easyjson6a975c40Encode1(out, in.Random)
	}
	{
		const prefix string = ",\"requestsLeft\":"
		out.RawString(prefix)
		out.Int64(int64(in.RequestsLeft))
	}
	out.RawByte('}')
}
func easyjson6a975c40Decode1(in *jlexer.Lexer, out *struct {
	CompletionTime string  `json:"completionTime"`
	Data           []int64 `json:"data"`
}) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "completionTime":
			out.CompletionTime = string(in.String())
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]int64, 0, 8)
					} else {
						out.Data = []int64{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v1 int64
					v1 = int64(in.Int64())
					out.Data = append(out.Data, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson6a975c40Encode1(out *jwriter.Writer, in struct {
	CompletionTime string  `json:"completionTime"`
	Data           []int64 `json:"data"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"completionTime\":"
		out.RawString(prefix[1:])
		out.String(string(in.CompletionTime))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		if in.Data == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Data {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson6a975c40DecodeJestRandomOrg1(in *jlexer.Lexer, out *RandomRequest) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int64(in.Int64())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		case "method":
			out.Method = string(in.String())
		case "params":
			easyjson6a975c40Decode2(in, &out.Params)
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
func easyjson6a975c40EncodeJestRandomOrg1(out *jwriter.Writer, in RandomRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		out.RawString(prefix)
		out.String(string(in.Jsonrpc))
	}
	{
		const prefix string = ",\"method\":"
		out.RawString(prefix)
		out.String(string(in.Method))
	}
	{
		const prefix string = ",\"params\":"
		out.RawString(prefix)
		easyjson6a975c40Encode2(out, in.Params)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RandomRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeJestRandomOrg1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RandomRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeJestRandomOrg1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RandomRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeJestRandomOrg1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RandomRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeJestRandomOrg1(l, v)
}
func easyjson6a975c40Decode2(in *jlexer.Lexer, out *struct {
	APIKey      string `json:"apiKey"`
	Max         int64  `json:"max"`
	Min         int64  `json:"min"`
	N           int64  `json:"n"`
	Replacement bool   `json:"replacement"`
}) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "apiKey":
			out.APIKey = string(in.String())
		case "max":
			out.Max = int64(in.Int64())
		case "min":
			out.Min = int64(in.Int64())
		case "n":
			out.N = int64(in.Int64())
		case "replacement":
			out.Replacement = bool(in.Bool())
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
func easyjson6a975c40Encode2(out *jwriter.Writer, in struct {
	APIKey      string `json:"apiKey"`
	Max         int64  `json:"max"`
	Min         int64  `json:"min"`
	N           int64  `json:"n"`
	Replacement bool   `json:"replacement"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"apiKey\":"
		out.RawString(prefix[1:])
		out.String(string(in.APIKey))
	}
	{
		const prefix string = ",\"max\":"
		out.RawString(prefix)
		out.Int64(int64(in.Max))
	}
	{
		const prefix string = ",\"min\":"
		out.RawString(prefix)
		out.Int64(int64(in.Min))
	}
	{
		const prefix string = ",\"n\":"
		out.RawString(prefix)
		out.Int64(int64(in.N))
	}
	{
		const prefix string = ",\"replacement\":"
		out.RawString(prefix)
		out.Bool(bool(in.Replacement))
	}
	out.RawByte('}')
}
