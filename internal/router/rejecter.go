package router

import (
	"gatekeeper-core/pkg/iso8583"
)

func (r *Router) CreateRejectMessage(req *iso8583.Message, reqMTI iso8583.MTI, responseCode iso8583.ResponseCode) *iso8583.Message {
	//resp := iso8583.NewMessage(reqMTI.ResponseMTI())
	//
	//resp.SetField(11, req.GetField(11)) // Systems Trace Audit Number (STAN)
	//resp.SetField(7, req.GetField(7))   // Transmission Date & Time
	//if req.GetField(3) != "" {
	//	resp.SetField(3, req.GetField(3)) // Processing Code обратно
	//}
	//if req.GetField(41) != "" {
	//	resp.SetField(41, req.GetField(41)) // Card Acceptor Terminal ID
	//}
	//
	//resp.SetField(39, string(responseCode))
	return nil
}
