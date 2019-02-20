package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"google.golang.org/grpc"

	"github.com/n0stack/n0stack/n0proto.go/deployment/v0"
)

func GetImage(c *cli.Context) error {
	endpoint := c.GlobalString("api-endpoint")
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Printf("[DEBUG] Connected to '%s'\n", endpoint)

	if c.NArg() == 0 {
		return listImage(conn)
	} else if c.NArg() == 1 {
		name := c.Args().Get(0)
		return getImage(name, conn)
	}

	return fmt.Errorf("set valid arguments")
}

func listImage(conn *grpc.ClientConn) error {
	cl := pdeployment.NewImageServiceClient(conn)
	res, err := cl.ListImages(context.Background(), &pdeployment.ListImagesRequest{})
	if err != nil {
		PrintGrpcError(err)
		return nil
	}

	marshaler.Marshal(os.Stdout, res)

	return nil
}

func getImage(name string, conn *grpc.ClientConn) error {
	cl := pdeployment.NewImageServiceClient(conn)
	res, err := cl.GetImage(context.Background(), &pdeployment.GetImageRequest{Name: name})
	if err != nil {
		PrintGrpcError(err)
		return nil
	}

	marshaler.Marshal(os.Stdout, res)

	return nil
}

func DeleteImage(c *cli.Context) error {
	endpoint := c.GlobalString("api-endpoint")
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Printf("[DEBUG] Connected to '%s'\n", endpoint)

	if c.NArg() == 1 {
		name := c.Args().Get(0)
		return deleteImage(name, conn)
	}

	return fmt.Errorf("set valid arguments")
}

func deleteImage(name string, conn *grpc.ClientConn) error {
	cl := pdeployment.NewImageServiceClient(conn)
	res, err := cl.DeleteImage(context.Background(), &pdeployment.DeleteImageRequest{Name: name})
	if err != nil {
		PrintGrpcError(err)
		return nil
	}

	marshaler.Marshal(os.Stdout, res)

	return nil
}