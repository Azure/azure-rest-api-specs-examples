package armcustomerlockbox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/customerlockbox/resource-manager/Microsoft.CustomerLockbox/preview/2018-02-28-preview/examples/Requests_List_FilterByStatus.json
func ExampleRequestsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcustomerlockbox.NewRequestsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		&armcustomerlockbox.RequestsClientListOptions{Filter: to.Ptr("properties/status eq 'Expired'")})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
