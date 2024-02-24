package frameproto

import (
	"sourcecode.social/reiver/go-opt"
)

// ActionBody represents the data sent to a Frame-Protocol's (i.e., Farcaster Frame's) application when it is POSTed to.
//
// Example usage:
//
//	p, err := io.ReadAll(request.Body)
//	if nil != err {
//		return err
//	}
//	
//	var actionBody frameproto.ActionBody
//	
//	err := json.Unmarshal(p, &actionBody)
//	if nil != err {
//		return err
//	}
type ActionBody struct {
	UntrustedData ActionBodyUntrustedData `json:"untrustedData"`
	TrustedData   ActionBodyTrustedData   `json:"trustedData"`
}

type ActionBodyUntrustedData struct {
	FID         opt.Optional[uint64] `json:"fid"`
	URL         opt.Optional[string] `json:"url"`
	MessageHash opt.Optional[string] `json:"messageHash"`
	Timestamp   opt.Optional[int64]  `json:"timestamp"`
	Network     opt.Optional[uint64] `json:"network"`
	ButtonIndex opt.Optional[uint64] `json:"buttonIndex"`
	InputText   opt.Optional[string] `json:"inputText"`
	CastID      ActionBodyCastID     `json:"castId"`
}

type ActionBodyTrustedData struct {
	MessageBytes opt.Optional[string] `json:"messageBytes"`
}

type ActionBodyCastID struct {
		FID  opt.Optional[uint64] `json:"fid"`
		Hash opt.Optional[string] `json:"hash"`
}
