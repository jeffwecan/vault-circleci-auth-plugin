package cciauth

import (
	"context"
	"math/rand"
	"fmt"
	"strconv"
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
			logical.ReadOperation:   b.pathGenerateNonceRead,
		},
	}
}

func (b *backend) pathGenerateNonceWrite(ctx context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	buildNum := d.Get("build_num").(int)

	nonce := make([]byte, 12)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	nonceString := hex.EncodeToString(nonce)

	//buildString := "nonce"+string(buildNum)
	buildString := fmt.Sprintf("nonce%s", strconv.Itoa(buildNum))

	entry, err := logical.StorageEntryJSON(buildString, nonceConfiguration{
		BuildNumber:	buildNum,
		Nonce:			nonceString,
	})

	if err != nil {
		return nil, err
	}

	if err := req.Storage.Put(ctx, entry); err != nil {
		return nil, err
	}

	resp := &logical.Response{
		Data: map[string]interface{}{
			"nonce": nonceString,
		},
	}

	return resp, nil
}

func (b *backend) pathGenerateNonceRead(ctx context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	buildNum := d.Get("build_num").(int)
	buildString := fmt.Sprintf("nonce%s", strconv.Itoa(buildNum))
	config, err := b.Nonce(ctx, req.Storage, buildString)
	if err != nil {
		return nil, err
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"Nonce":       	config.Nonce,
			"build_num":  	config.BuildNumber,
		},
	}, nil
}

// Config reads the nonce object out of the provided Storage.
func (b *backend) Nonce(ctx context.Context, s logical.Storage, buildNum string) (*nonceConfiguration, error) {
	entry, err := s.Get(ctx, buildNum)
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
	Nonce			string
}
