package main

import (
	"fmt"
	"log"

	"beeswax/bid"

	"github.com/golang/protobuf/proto"
	"github.com/valyala/fasthttp"
)

func stringPtr(s string) *string {
	return &s
}

func contains(item []string, search string) bool {
	for _, e := range item {
		if search == e {
			return true
		}
	}
	return false
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	bidAgentRequest := bid.BidAgentRequest{}
	if err := proto.Unmarshal(ctx.PostBody(), &bidAgentRequest); err != nil {
		fmt.Println(err)
	}
	bidRequest := bidAgentRequest.GetBidRequest()

	// App or site
	// IAB2-3 (Buying/Selling Cars)=> $2.1 or $2.0
	var bidPrice uint64
	if bidRequest.App != nil {
		if contains(bidRequest.GetApp().Cat, "IAB2-3") {
			bidPrice = 2100
		}
	} else if bidRequest.Site != nil {
		if contains(bidRequest.GetSite().Cat, "IAB2-3") {
			bidPrice = 2000
		}
	} else {
		// we are not interested for a bid
		ctx.SetStatusCode(fasthttp.StatusNoContent)
		// TODO fasthttp reuse the ctx instance for performance raison (zero allocation)
		// we should CLEAN it before returning !!!
		return
	}

	bidResponse := bid.BidAgentResponse{
		Bids: []*bid.BidAgentResponse_Bid{
			&bid.BidAgentResponse_Bid{
				BidPriceMicros: &bidPrice,
				LineItemId:     bidAgentRequest.GetAdcandidates()[0].LineItemId,      // probably a bad idea to bid on something we don't know :-)
				CreativeId:     &bidAgentRequest.GetAdcandidates()[0].CreativeIds[0], // same...
			},
		},
	}
	respByte, err := proto.Marshal(&bidResponse)
	if err != nil {
		log.Fatalln(err)
	}
	ctx.SetBody(respByte)
}

func main() {
	fasthttp.ListenAndServe(":3001", fastHTTPHandler)
}
