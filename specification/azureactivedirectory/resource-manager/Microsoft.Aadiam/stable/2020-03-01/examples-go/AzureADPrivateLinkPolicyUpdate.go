package armaad_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/examples/AzureADPrivateLinkPolicyUpdate.json
func ExamplePrivateLinkForAzureAdClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armaad.NewPrivateLinkForAzureAdClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"rg1",
		"ddb1",
		&armaad.PrivateLinkForAzureAdClientUpdateOptions{PrivateLinkPolicy: &armaad.PrivateLinkPolicyUpdateParameter{
			Tags: map[string]*string{},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
