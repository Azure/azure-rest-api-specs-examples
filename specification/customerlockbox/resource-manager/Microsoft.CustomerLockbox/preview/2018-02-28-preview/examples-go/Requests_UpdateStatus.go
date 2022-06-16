package armcustomerlockbox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/customerlockbox/resource-manager/Microsoft.CustomerLockbox/preview/2018-02-28-preview/examples/Requests_UpdateStatus.json
func ExampleRequestsClient_UpdateStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcustomerlockbox.NewRequestsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateStatus(ctx,
		"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		armcustomerlockbox.Approval{
			Reason: to.Ptr("Customer approve"),
			Status: to.Ptr(armcustomerlockbox.StatusApprove),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
