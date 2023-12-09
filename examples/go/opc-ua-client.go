package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/ua"
)

func main() {

	s := gocron.NewScheduler(time.UTC)

	var (
		endpoint = flag.String("endpoint", "opc.tcp://10.0.0.130:4840", "OPC UA Endpoint URL")
		nodeID   = flag.String("node", "ns=4;s=|var|Edge Controller PIO3.Application.PLC_PRG.iCnt", "NodeID to read")
	)

	flag.BoolVar(&debug.Enable, "debug", false, "enable debug logging")
	flag.Parse()
	log.SetFlags(0)

	eps, err := opcua.GetEndpoints(context.Background(), *endpoint)
	if err != nil {
		log.Fatal(err)
	}

	for _, ep := range eps {
		log.Println(ep.EndpointURL, ep.SecurityPolicyURI, ep.SecurityMode)
		log.Println("Wago device found!")
	}

	ctx := context.Background()

	c := opcua.NewClient(*endpoint, opcua.SecurityMode(ua.MessageSecurityModeNone))
	if err := c.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer c.CloseWithContext(ctx)

	id, err := ua.ParseNodeID(*nodeID)
	if err != nil {
		log.Fatalf("invalid node id: %v", err)
	}

	req := &ua.ReadRequest{
		MaxAge: 2000,
		NodesToRead: []*ua.ReadValueID{
			{NodeID: id},
		},
		TimestampsToReturn: ua.TimestampsToReturnBoth,
	}

	s.Every(1).Second().Do(func() {

		resp, err := c.ReadWithContext(ctx, req)
		if err != nil {
			log.Fatalf("Read failed: %s", err)
		}

		if resp.Results[0].Status != ua.StatusOK {
			log.Fatalf("Status not OK: %v", resp.Results[0].Status)
		}
		log.Println("iCnt: ", resp.Results[0].Value.Value())

	})

	s.StartBlocking()

}
