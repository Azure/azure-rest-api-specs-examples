package armmixedreality_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/RegenerateKey.json
func ExampleObjectAnchorsAccountsClient_RegenerateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmixedreality.NewObjectAnchorsAccountsClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.RegenerateKeys(ctx,
		"MyResourceGroup",
		"MyAccount",
		armmixedreality.AccountKeyRegenerateRequest{
			Serial: to.Ptr(armmixedreality.SerialPrimary),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
