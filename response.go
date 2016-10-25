package enable

import (
	"github.com/emersion/go-imap"
)

// An ENABLE response, defined in RFC 5161 section 3.2.
type Response struct {
	Capabilities []string
}

func (r *Response) HandleFrom(hdlr imap.RespHandler) error {
	r.Capabilities = nil

	for h := range hdlr {
		fields, ok := h.AcceptNamedResp(responseName)
		if !ok {
			continue
		}

		if caps, err := imap.ParseStringList(fields); err != nil {
			return err
		} else {
			r.Capabilities = append(r.Capabilities, caps...)
		}
	}

	return nil
}

func (r *Response) WriteTo(w *imap.Writer) error {
	fields := []interface{}{responseName}
	fields = append(fields, imap.FormatStringList(r.Capabilities)...)

	res := &imap.Resp{Fields: fields}
	return res.WriteTo(w)
}
