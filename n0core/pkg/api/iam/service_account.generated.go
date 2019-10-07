// Code generated by "stdapi"; DO NOT EDIT.

package iam

import (
	"context"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"

	piam "n0st.ac/n0stack/iam/v1alpha"
	stdapi "n0st.ac/n0stack/n0core/pkg/api/stdapi"
	"n0st.ac/n0stack/n0core/pkg/datastore"
	grpcutil "n0st.ac/n0stack/n0core/pkg/util/grpc"
)

func GetServiceAccount(ctx context.Context, ds datastore.Datastore, name string) (*piam.ServiceAccount, int64, error) {
	resourse := &piam.ServiceAccount{}
	version, err := ds.Get(ctx, name, resourse)
	if err != nil {
		if datastore.IsNotFound(err) {
			return nil, 0, grpcutil.Errorf(codes.NotFound, err.Error())
		}

		return nil, 0, grpcutil.Errorf(codes.Internal, "failed to get User %s from db: err='%s'", name, err.Error())
	}

	return resourse, version, nil
}

func ListServiceAccounts(ctx context.Context, req *piam.ListServiceAccountsRequest, ds datastore.Datastore) (*piam.ListServiceAccountsResponse, error) {
	res := &piam.ListServiceAccountsResponse{}
	f := func(s int) []proto.Message {
		res.ServiceAccounts = make([]*piam.ServiceAccount, s)
		for i := range res.ServiceAccounts {
			res.ServiceAccounts[i] = &piam.ServiceAccount{}
		}

		m := make([]proto.Message, s)
		for i, v := range res.ServiceAccounts {
			m[i] = v
		}

		return m
	}

	if err := ds.List(ctx, f); err != nil {
		return nil, grpcutil.Errorf(codes.Internal, "Failed to list from db, please retry or contact for the administrator of this cluster")
	}
	if len(res.ServiceAccounts) == 0 {
		return nil, grpcutil.Errorf(codes.NotFound, "")
	}

	return res, nil
}

func DeleteServiceAccount(ctx context.Context, ds datastore.Datastore, name string, version int64) error {
	if err := ds.Delete(ctx, name, version); err != nil {
		return stdapi.DatastoreDeleteError(err, "ServiceAccount", name)
	}

	return nil
}

func ApplyServiceAccount(ctx context.Context, ds datastore.Datastore, resource *piam.ServiceAccount, version int64) (int64, error) {
	version, err := ds.Apply(ctx, resource.Name, resource, version)
	if err != nil {
		return 0, stdapi.DatastoreApplyError(err, "ServiceAccount", resource.Name)
	}

	return version, nil
}
