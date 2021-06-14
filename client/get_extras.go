package client

import (
	"context"
	"errors"
)

func (s *Client) GetGenesisHash(ctx context.Context) (string, error) {
	res := struct {
		GeneralResponse
		Result string `json:"result"`
	}{}
	err := s.request(ctx, "getGenesisHash", []interface{}{}, &res)
	if err != nil {
		return "", err
	}
	if res.Error != (ErrorResponse{}) {
		return "", errors.New(res.Error.Message)
	}
	return res.Result, nil
}
func (s *Client) GetFirstAvailableBlock(ctx context.Context) (uint64, error) {
	res := struct {
		GeneralResponse
		Result uint64 `json:"result"`
	}{}
	err := s.request(ctx, "getFirstAvailableBlock", []interface{}{}, &res)
	if err != nil {
		return 0, err
	}
	if res.Error != (ErrorResponse{}) {
		return 0, errors.New(res.Error.Message)
	}
	return res.Result, nil
}

func (s *Client) GetSlot(ctx context.Context) (uint64, error) {
	res := struct {
		GeneralResponse
		Result uint64 `json:"result"`
	}{}
	err := s.request(ctx, "getSlot", []interface{}{}, &res)
	if err != nil {
		return 0, err
	}
	if res.Error != (ErrorResponse{}) {
		return 0, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
