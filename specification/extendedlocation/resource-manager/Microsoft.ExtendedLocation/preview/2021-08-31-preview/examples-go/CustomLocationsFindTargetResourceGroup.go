package armextendedlocation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fb9c8e2ca33e9723c2b2fc849f627329067feb54/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/CustomLocationsFindTargetResourceGroup.json
func ExampleCustomLocationsClient_FindTargetResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomLocationsClient().FindTargetResourceGroup(ctx, "testresourcegroup", "customLocation01", armextendedlocation.CustomLocationFindTargetResourceGroupProperties{
		Labels: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomLocationFindTargetResourceGroupResult = armextendedlocation.CustomLocationFindTargetResourceGroupResult{
	// 	MatchedResourceSyncRule: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ExtendedLocation/customLocations/tesla-customlocation-1-chumpishness/resourceSyncRules/test-resource-sync-rule"),
	// 	TargetResourceGroup: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/target-resource-group/"),
	// }
}
