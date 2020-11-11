package cciauth

import (
	"context"
	"math/rand"
	"fmt"
	"encoding/hex"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func (b *backend) pathGenerateNonce() *framework.Path {
	return &framework.Path{
		Pattern: "nonce",
		Fields: map[string]*framework.FieldSchema{
			"build_num": {
				Type:        framework.TypeInt,
				Description: "The number of the current build.",
			},
		},
		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.UpdateOperation: b.pathGenerateNonceWrite,
		},
	}
}

func (b *backend) pathGenerateNonceWrite(ctx context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	buildNum := d.Get("build_num").(int)

	nonce := make([]byte, 12)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	entry, err := logical.StorageEntryJSON("nonce", nonceConfiguration{
		BuildNumber:	buildNum,
		Nonce:			nonce,
	})

	if err != nil {
		return nil, err
	}

	if err := req.Storage.Put(ctx, entry); err != nil {
		return nil, err
	}

	// print nonce to stdout - figure out how to print this as part of the response
	fmt.Println("Nonce is:", hex.EncodeToString(nonce))

	return nil, nil
}

// Config reads the nonce object out of the provided Storage.
func (b *backend) Nonce(ctx context.Context, s logical.Storage) (*nonceConfiguration, error) {
	entry, err := s.Get(ctx, "nonce")
	if err != nil {
		return nil, err
	}

	var result nonceConfiguration
	if entry != nil {
		if err := entry.DecodeJSON(&result); err != nil {
			return nil, fmt.Errorf("error reading configuration: %s", err)
		}
	}

	return &result, nil
}

type nonceConfiguration struct {
	BuildNumber      int        `json:"build_number"`
	Nonce			[]byte
}
