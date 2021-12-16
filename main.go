package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hellofresh/storev2/market"
	as "github.com/igor-tonkopryadchenko/store/grpc/service"
	"google.golang.org/grpc"
)

const grpcServerEndpoint = "localhost:9090"

type AddonsServer struct {
	as.UnimplementedAddonsServiceServer
}

func (s AddonsServer) ListSections(ctx context.Context, req *as.ListSectionsRequest) (*as.ListSectionsResponse, error) {
	h := market.Section{Handle: "some"}
	ss := []*market.Section{&h}

	resp := as.ListSectionsResponse{Sections: ss}

	fmt.Println("req", req)
	fmt.Println("req valid?", req.ValidateAll())

	w := market.YearWeek{Year: 1998, Week: 44}
	fmt.Println("week valid?", w.HfString(), w.ValidateAll())

	l := market.Locale{Lang: 7, Country: market.Country_DE}

	fmt.Println("locale valid?", &l, l.ValidateAll())

	return &resp, nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// opts := []grpc.DialOption{insecure.NewCredentials()}
	err := as.RegisterAddonsServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	s := AddonsServer{}
	as.RegisterAddonsServiceHandlerServer(ctx, mux, s)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	fmt.Println("Starting proxy on :8081...")
	return http.ListenAndServe(":8081", mux)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
