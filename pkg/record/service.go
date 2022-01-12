package record

import (
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const serviceName = "Record"

type Service struct {
	c *client.Client
}

func New(cnf client.Config) (*Service, error) {
	c, err := client.NewClient(cnf)

	if err != nil {
		return nil, helper.ServiceConfigError(serviceName, err)
	}

	return &Service{c}, nil
}

func Get(c *client.Client) (*Service, error) {
	if c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	return &Service{c}, nil
}

func (s *Service) CreateRecord(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPost, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, helper.CreateError(serviceName, rrSetKey.ID(), err)
	}

	return res, nil
}

func (s *Service) UpdateRecord(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPut, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, helper.UpdateError(serviceName, rrSetKey.ID(), err)
	}

	return res, nil
}

func (s *Service) PartialUpdateRecord(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodPatch, rrSetKey.URI(), rrSet, target)

	if err != nil {
		return nil, helper.PartialUpdateError(serviceName, rrSetKey.ID(), err)
	}

	return res, nil
}

func (s *Service) ReadRecord(rrSetKey *rrset.RRSetKey) (*http.Response, *rrset.ResponseList, error) {
	target := client.Target(&rrset.ResponseList{})

	if s.c == nil {
		return nil, nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, rrSetKey.URI(), nil, target)

	if err != nil {
		return nil, nil, helper.ReadError(serviceName, rrSetKey.ID(), err)
	}

	rrsetList := target.Data.(*rrset.ResponseList)

	return res, rrsetList, nil
}

func (s *Service) DeleteRecord(rrSetKey *rrset.RRSetKey) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, helper.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodDelete, rrSetKey.URI(), nil, target)

	if err != nil {
		return nil, helper.DeleteError(serviceName, rrSetKey.ID(), err)
	}

	return res, nil
}
