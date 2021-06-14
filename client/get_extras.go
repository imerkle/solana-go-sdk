package client

import (
	"context"
)

func (s *Client) GetGenesisHash(ctx context.Context) (string, error) {
	var res string
	_ = s.request(ctx, "getGenesisHash", []interface{}{}, &res)

	return res, nil
}
func (s *Client) GetFirstAvailableBlock(ctx context.Context) (uint64, error) {
	var res uint64
	_ = s.request(ctx, "getFirstAvailableBlock", []interface{}{}, &res)

	return res, nil
}
func (s *Client) GetSlot(ctx context.Context) (uint64, error) {
	var res uint64
	_ = s.request(ctx, "getSlot", []interface{}{}, &res)

	return res, nil
}
