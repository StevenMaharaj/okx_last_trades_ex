package main


type OkxSubReq struct {
	Op   string `json:"op"`
	OkxSubReqArgs []OkxSubReqArg `json:"args"`
}
type OkxSubReqArg struct {
	Channel string `json:"channel"`
	InstID  string `json:"instId"`
}


type Trade struct {
  price float64
  volume float64
}

type OkxSubResp struct {
	Event  string `json:"event"`
	OkxSubReqArg   OkxSubReqArg   `json:"args"`
	ConnID string `json:"connId"`
}


type OkxRespFail struct {
	Event  string `json:"event"`
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	ConnID string `json:"connId"`
}

type LastTradePush struct {
	Arg struct {
		Channel string `json:"channel"`
		InstID  string `json:"instId"`
	} `json:"arg"`
	Data []struct {
		InstID  string `json:"instId"`
		TradeID string `json:"tradeId"`
		Px      string `json:"px"`
		Sz      string `json:"sz"`
		Side    string `json:"side"`
		Ts      string `json:"ts"`
		Count   string `json:"count"`
	} `json:"data"`
}
