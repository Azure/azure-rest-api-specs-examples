package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/AttachedNetworks_GetByDevCenter.json
func ExampleAttachedNetworksClient_GetByDevCenter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevcenter.NewAttachedNetworksClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetByDevCenter(ctx,
		"rg1",
		"Contoso",
		"network-uswest3",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
