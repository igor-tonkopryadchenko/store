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

func (s AddonsServer) ListSections(ctx context.Context, req *as.ListSectionsRequest) (*as.ListSectionsResponse, error) {
	h := domain.Section{Handle: "some"}
	ss := []*domain.Section{&h}

	resp := as.ListSectionsResponse{Sections: ss}

	fmt.Println("req", req)
	fmt.Println("req valid?", req.ValidateAll())

	w := domain.YearWeek{Year: 1998, Week: 44}
	fmt.Println("week valid?", w.String(), w.ValidateAll())

	l := domain.Locale{Lang: 7, Country: domain.Country_DE}

	fmt.Println("locale valid?", &l, l.ValidateAll())

	return &resp, nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	gwmux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	// opts := []grpc.DialOption{insecure.NewCredentials()}
	err := as.RegisterAddonsServiceHandlerFromEndpoint(ctx, gwmux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	mux.Handle("/", gwmux)

	mux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(">> path: ", r.URL.Path)
		fmt.Println(">> filepath: ", r.URL.Path[1:])

		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// sh := http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger/")))
	fs := http.FileServer(http.Dir("./swagger"))
	mux.Handle("/swagger2/", http.StripPrefix("/swagger2/", fs))
	// r.PathPrefix("/swaggerui/").Handler(sh)

	s := AddonsServer{}
	as.RegisterAddonsServiceHandlerServer(ctx, gwmux, s)

	gs := grpc.NewServer()
	as.RegisterAddonsServiceServer(gs, s)

	l, err := net.Listen("tcp", "127.0.0.1:9090")

	if err != nil {
		return err
	}

	gs.Serve(l)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	fmt.Println("Starting proxy on :8081...")
	return http.ListenAndServe(":8081", mux)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
