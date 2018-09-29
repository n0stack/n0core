package main

import (
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/n0stack/n0core/pkg/api/pool/node"
	"github.com/n0stack/n0core/pkg/api/provisioning"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Serve(ctx *cli.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ctx.String("bind-address"), ctx.Int("bind-port")))
	if err != nil {
		return err
	}

	bvm := filepath.Join(ctx.String("base-directory"), "virtual_machine")
	vma, err := provisioning.CreateVirtualMachineAgentAPI(bvm)
	if err != nil {
		return err
	}

	bv := filepath.Join(ctx.String("base-directory"), "volume")
	va, err := provisioning.CreateBlockStorageAgentAPI(bv)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	provisioning.RegisterBlockStorageAgentServiceServer(s, va)
	provisioning.RegisterVirtualMachineAgentServiceServer(s, vma)
	reflection.Register(s)

	if err := node.RegisterNodeToAPI(ctx.String("name"), ctx.String("advertise-address"), ctx.String("node-api-endpoint")); err != nil {
		return err
	}

	log.Printf("[INFO] Started API")
	return s.Serve(lis)
}
