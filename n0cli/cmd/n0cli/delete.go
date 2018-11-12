package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/n0stack/n0stack/n0proto.go/deployment/v0"
	"github.com/n0stack/n0stack/n0proto.go/pool/v0"
	"github.com/n0stack/n0stack/n0proto.go/provisioning/v0"

	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func Delete(ctx *cli.Context) error {
	if ctx.NArg() == 2 {
		return delete(ctx)
	}

	return fmt.Errorf("set valid arguments")
}

func delete(ctx *cli.Context) error {
	resourceType := ctx.Args().Get(0)
	resourceName := ctx.Args().Get(1)

	endpoint := ctx.GlobalString("api-endpoint")
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Printf("[DEBUG] Connected to '%s'\n", endpoint)

	switch resourceType {
	case "node":
		cl := ppool.NewNodeServiceClient(conn)
		res, err := cl.DeleteNode(context.Background(), &ppool.DeleteNodeRequest{Name: resourceName})
		if err != nil {
			fmt.Printf("got error response: %s\n", err.Error())
			return nil
		}

		marshaler.Marshal(os.Stdout, res)

	case "network":
		cl := ppool.NewNetworkServiceClient(conn)
		res, err := cl.DeleteNetwork(context.Background(), &ppool.DeleteNetworkRequest{Name: resourceName})
		if err != nil {
			fmt.Printf("got error response: %s\n", err.Error())
			return nil
		}

		marshaler.Marshal(os.Stdout, res)

	case "block_storage":
		cl := pprovisioning.NewBlockStorageServiceClient(conn)
		res, err := cl.DeleteBlockStorage(context.Background(), &pprovisioning.DeleteBlockStorageRequest{Name: resourceName})
		if err != nil {
			fmt.Printf("got error response: %s\n", err.Error())
			return nil
		}

		marshaler.Marshal(os.Stdout, res)

	case "virtual_machine":
		cl := pprovisioning.NewVirtualMachineServiceClient(conn)
		res, err := cl.DeleteVirtualMachine(context.Background(), &pprovisioning.DeleteVirtualMachineRequest{Name: resourceName})
		if err != nil {
			fmt.Printf("got error response: %s\n", err.Error())
			return nil
		}

		marshaler.Marshal(os.Stdout, res)

	case "image":
		cl := pdeployment.NewImageServiceClient(conn)
		res, err := cl.DeleteImage(context.Background(), &pdeployment.DeleteImageRequest{Name: resourceName})
		if err != nil {
			fmt.Printf("got error response: %s\n", err.Error())
			return nil
		}

		marshaler.Marshal(os.Stdout, res)

	case "flavor":
		cl := pdeployment.NewFlavorServiceClient(conn)
		res, err := cl.DeleteFlavor(context.Background(), &pdeployment.DeleteFlavorRequest{Name: resourceName})
		if err != nil {
			fmt.Printf("got error response: %s\n", err.Error())
			return nil
		}

		marshaler.Marshal(os.Stdout, res)

	default:
		return fmt.Errorf("resource type '%s' is not existing\n", resourceType)
	}

	return nil
}