package fakes

import (
	"github.com/SAP/cf-mta-plugin/clients/cfrestclient"
	"github.com/SAP/cf-mta-plugin/clients/models"
)

// TODO: use counterfeiter if the client becomes more suffisticated

type FakeCloudFoundryClient struct {
	domains []models.SharedDomain
	err     error
}

func NewFakeCloudFoundryClient(domains []models.SharedDomain, err error) cfrestclient.CloudFoundryOperationsExtended {
	return FakeCloudFoundryClient{domains: domains, err: err}
}

func (f FakeCloudFoundryClient) GetSharedDomains() ([]models.SharedDomain, error) {
	return f.domains, f.err
}
