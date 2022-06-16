package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2020-05-01/examples/AvailabilityStatuses_ListByResourceGroup.json
func ExampleAvailabilityStatusesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresourcehealth.NewAvailabilityStatusesClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("resourceGroupName",
		&armresourcehealth.AvailabilityStatusesClientListByResourceGroupOptions{Filter: nil,
			Expand: to.Ptr("recommendedactions"),
		})
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
