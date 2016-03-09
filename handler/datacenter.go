package handler

import (
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/kubermatic/api"
	"github.com/kubermatic/api/provider"
	"golang.org/x/net/context"
)

func datacentersEndpoint(
	dcs map[string]provider.DatacenterMeta,
	kps map[string]provider.KubernetesProvider,
	cps map[string]provider.CloudProvider,
) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		adcs := make([]api.Datacenter, 0, len(kps))
		for dcName := range dcs {
			_, kpFound := kps[dcName]
			dc := dcs[dcName]

			spec, err := apiSpec(&dc)
			if err != nil {
				glog.Errorf("api spec error in dc %q: %v", dcName, err)
				continue
			}

			adc := api.Datacenter{
				Metadata: api.Metadata{
					Name:     dcName,
					Revision: "1",
				},
				Spec: *spec,
				Seed: kpFound,
			}
			adcs = append(adcs, adc)
		}

		return adcs, nil
	}
}

func datacenterEndpoint(
	dcs map[string]provider.DatacenterMeta,
	kps map[string]provider.KubernetesProvider,
	cps map[string]provider.CloudProvider,
) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dcReq)

		dc, found := dcs[req.dc]
		if !found {
			return nil, NewNotFound("datacenter", req.dc)
		}

		_, kpFound := kps[req.dc]

		spec, err := apiSpec(&dc)
		if err != nil {
			return nil, fmt.Errorf("api spec error in dc %q: %v", req.dc, err)
		}

		return &api.Datacenter{
			Metadata: api.Metadata{
				Name:     req.dc,
				Revision: "1",
			},
			Spec: *spec,
			Seed: kpFound,
		}, nil
	}
}

type dcsReq struct {
}

func decodeDatacentersReq(r *http.Request) (interface{}, error) {
	return dcsReq{}, nil
}

func decodeDatacenterReq(r *http.Request) (interface{}, error) {
	return dcReq{
		dc: mux.Vars(r)["dc"],
	}, nil
}

func apiSpec(dc *provider.DatacenterMeta) (*api.DatacenterSpec, error) {
	p, err := provider.DatacenterCloudProviderName(&dc.Spec)
	if err != nil {
		return nil, err
	}
	spec := &api.DatacenterSpec{
		Location: dc.Location,
		Country:  dc.Country,
		Provider: p,
	}

	switch {
	case dc.Spec.Digitalocean != nil:
		spec.Digitalocean = &api.DigitialoceanDatacenterSpec{
			Region: dc.Spec.Digitalocean.Region,
		}
	case dc.Spec.BringYourOwn != nil:
		spec.BringYourOwn = &api.BringYourOwnDatacenterSpec{}
	}

	return spec, nil
}
