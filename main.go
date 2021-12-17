package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/igor-tonkopryadchenko/store/domain"
	as "github.com/igor-tonkopryadchenko/store/service"
	"google.golang.org/grpc"
)

const grpcServerEndpoint = "localhost:9090"

type AddonsServer struct {
	as.UnimplementedAddonsServiceServer
}

func (s AddonsServer) UpdateSection(ctx context.Context, in *as.UpdateSectionRequest) (*as.NoResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, err
	}

	fmt.Println(">valid request: ", in)

	return &as.NoResponse{}, nil
}

func (s AddonsServer) ListSections(ctx context.Context, req *as.ListSectionsRequest) (*as.ListSectionsResponse, error) {
	h := domain.Section{Handle: "some"}
	ss := []*domain.Section{&h}

	resp := as.ListSectionsResponse{Sections: ss}

	fmt.Println("req", req)

	if err := req.ValidateAll(); err != nil {
		fmt.Println("req invalid:", err)
		return nil, err
	}

	w := domain.YearWeek{Year: 1998, Week: 44}
	fmt.Println("week valid?", w.String(), w.ValidateAll())

	l := domain.Locale{Lang: 7, Country: domain.Country_DE}

	fmt.Println("locale valid?", &l, l.ValidateAll())

	return &resp, nil
}

func runHttpProxy(ctx context.Context, addr string) error {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	gwmux := runtime.NewServeMux()
	mux.Handle("/", gwmux)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := as.RegisterAddonsServiceHandlerFromEndpoint(ctx, gwmux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	fmt.Println("Starting http proxy on ", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println(">> oi!", err)
		return err
	}

	return nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s := AddonsServer{}

	gs := grpc.NewServer()
	as.RegisterAddonsServiceServer(gs, s)

	l, err := net.Listen("tcp", grpcServerEndpoint)

	if err != nil {
		return err
	}

	go runHttpProxy(ctx, ":8081")

	fmt.Println("Starting grpc on ", grpcServerEndpoint)
	return gs.Serve(l)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
