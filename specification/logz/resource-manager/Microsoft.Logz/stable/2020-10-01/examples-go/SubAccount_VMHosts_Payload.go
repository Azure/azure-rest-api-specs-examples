package armlogz_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logz/armlogz"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_VMHosts_Payload.json
func ExampleSubAccountClient_VMHostPayload() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armlogz.NewSubAccountClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.VMHostPayload(ctx,
		"myResourceGroup",
		"myMonitor",
		"SubAccount1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
