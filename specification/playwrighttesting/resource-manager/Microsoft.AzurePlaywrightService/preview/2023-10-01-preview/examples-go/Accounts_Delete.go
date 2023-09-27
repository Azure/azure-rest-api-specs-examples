package armplaywrighttesting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/playwrighttesting/armplaywrighttesting"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3eb9ec8e9c8f717c6b461c4c0f49a4662fb948fd/specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/examples/Accounts_Delete.json
func ExampleAccountsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplaywrighttesting.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginDelete(ctx, "dummyrg", "myPlaywrightAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
