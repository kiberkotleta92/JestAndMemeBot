// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package vk

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

func easyjson6a975c40DecodeJestVk(in *jlexer.Lexer, out *VKResponse) {
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
		case "response":
			easyjson6a975c40Decode(in, &out.Response)
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
func easyjson6a975c40EncodeJestVk(out *jwriter.Writer, in VKResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"response\":"
		out.RawString(prefix[1:])
		easyjson6a975c40Encode(out, in.Response)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VKResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeJestVk(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VKResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeJestVk(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VKResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeJestVk(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VKResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeJestVk(l, v)
}
func easyjson6a975c40Decode(in *jlexer.Lexer, out *struct {
	Count int64 `json:"count"`
	Items []struct {
		AlbumID int64 `json:"album_id"`
		Date    int64 `json:"date"`
		ID      int64 `json:"id"`
		OwnerID int64 `json:"owner_id"`
		PostID  int64 `json:"post_id"`
		Sizes   []struct {
			Height int64  `json:"height"`
			Type   string `json:"type"`
			URL    string `json:"url"`
			Width  int64  `json:"width"`
		} `json:"sizes"`
		Text   string `json:"text"`
		UserID int64  `json:"user_id"`
	} `json:"items"`
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
		case "count":
			out.Count = int64(in.Int64())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]struct {
							AlbumID int64 `json:"album_id"`
							Date    int64 `json:"date"`
							ID      int64 `json:"id"`
							OwnerID int64 `json:"owner_id"`
							PostID  int64 `json:"post_id"`
							Sizes   []struct {
								Height int64  `json:"height"`
								Type   string `json:"type"`
								URL    string `json:"url"`
								Width  int64  `json:"width"`
							} `json:"sizes"`
							Text   string `json:"text"`
							UserID int64  `json:"user_id"`
						}, 0, 1)
					} else {
						out.Items = []struct {
							AlbumID int64 `json:"album_id"`
							Date    int64 `json:"date"`
							ID      int64 `json:"id"`
							OwnerID int64 `json:"owner_id"`
							PostID  int64 `json:"post_id"`
							Sizes   []struct {
								Height int64  `json:"height"`
								Type   string `json:"type"`
								URL    string `json:"url"`
								Width  int64  `json:"width"`
							} `json:"sizes"`
							Text   string `json:"text"`
							UserID int64  `json:"user_id"`
						}{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v1 struct {
						AlbumID int64 `json:"album_id"`
						Date    int64 `json:"date"`
						ID      int64 `json:"id"`
						OwnerID int64 `json:"owner_id"`
						PostID  int64 `json:"post_id"`
						Sizes   []struct {
							Height int64  `json:"height"`
							Type   string `json:"type"`
							URL    string `json:"url"`
							Width  int64  `json:"width"`
						} `json:"sizes"`
						Text   string `json:"text"`
						UserID int64  `json:"user_id"`
					}
					easyjson6a975c40Decode1(in, &v1)
					out.Items = append(out.Items, v1)
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
func easyjson6a975c40Encode(out *jwriter.Writer, in struct {
	Count int64 `json:"count"`
	Items []struct {
		AlbumID int64 `json:"album_id"`
		Date    int64 `json:"date"`
		ID      int64 `json:"id"`
		OwnerID int64 `json:"owner_id"`
		PostID  int64 `json:"post_id"`
		Sizes   []struct {
			Height int64  `json:"height"`
			Type   string `json:"type"`
			URL    string `json:"url"`
			Width  int64  `json:"width"`
		} `json:"sizes"`
		Text   string `json:"text"`
		UserID int64  `json:"user_id"`
	} `json:"items"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"count\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Count))
	}
	{
		const prefix string = ",\"items\":"
		out.RawString(prefix)
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Items {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjson6a975c40Encode1(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson6a975c40Decode1(in *jlexer.Lexer, out *struct {
	AlbumID int64 `json:"album_id"`
	Date    int64 `json:"date"`
	ID      int64 `json:"id"`
	OwnerID int64 `json:"owner_id"`
	PostID  int64 `json:"post_id"`
	Sizes   []struct {
		Height int64  `json:"height"`
		Type   string `json:"type"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"sizes"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
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
		case "album_id":
			out.AlbumID = int64(in.Int64())
		case "date":
			out.Date = int64(in.Int64())
		case "id":
			out.ID = int64(in.Int64())
		case "owner_id":
			out.OwnerID = int64(in.Int64())
		case "post_id":
			out.PostID = int64(in.Int64())
		case "sizes":
			if in.IsNull() {
				in.Skip()
				out.Sizes = nil
			} else {
				in.Delim('[')
				if out.Sizes == nil {
					if !in.IsDelim(']') {
						out.Sizes = make([]struct {
							Height int64  `json:"height"`
							Type   string `json:"type"`
							URL    string `json:"url"`
							Width  int64  `json:"width"`
						}, 0, 1)
					} else {
						out.Sizes = []struct {
							Height int64  `json:"height"`
							Type   string `json:"type"`
							URL    string `json:"url"`
							Width  int64  `json:"width"`
						}{}
					}
				} else {
					out.Sizes = (out.Sizes)[:0]
				}
				for !in.IsDelim(']') {
					var v4 struct {
						Height int64  `json:"height"`
						Type   string `json:"type"`
						URL    string `json:"url"`
						Width  int64  `json:"width"`
					}
					easyjson6a975c40Decode2(in, &v4)
					out.Sizes = append(out.Sizes, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "text":
			out.Text = string(in.String())
		case "user_id":
			out.UserID = int64(in.Int64())
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
	AlbumID int64 `json:"album_id"`
	Date    int64 `json:"date"`
	ID      int64 `json:"id"`
	OwnerID int64 `json:"owner_id"`
	PostID  int64 `json:"post_id"`
	Sizes   []struct {
		Height int64  `json:"height"`
		Type   string `json:"type"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"sizes"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"album_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.AlbumID))
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.Int64(int64(in.Date))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"owner_id\":"
		out.RawString(prefix)
		out.Int64(int64(in.OwnerID))
	}
	{
		const prefix string = ",\"post_id\":"
		out.RawString(prefix)
		out.Int64(int64(in.PostID))
	}
	{
		const prefix string = ",\"sizes\":"
		out.RawString(prefix)
		if in.Sizes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Sizes {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson6a975c40Encode2(out, v6)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix)
		out.Int64(int64(in.UserID))
	}
	out.RawByte('}')
}
func easyjson6a975c40Decode2(in *jlexer.Lexer, out *struct {
	Height int64  `json:"height"`
	Type   string `json:"type"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
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
		case "height":
			out.Height = int64(in.Int64())
		case "type":
			out.Type = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "width":
			out.Width = int64(in.Int64())
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
	Height int64  `json:"height"`
	Type   string `json:"type"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"height\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Height))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	{
		const prefix string = ",\"width\":"
		out.RawString(prefix)
		out.Int64(int64(in.Width))
	}
	out.RawByte('}')
}