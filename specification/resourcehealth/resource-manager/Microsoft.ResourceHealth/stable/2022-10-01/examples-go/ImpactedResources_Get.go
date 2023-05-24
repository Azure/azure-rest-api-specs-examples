package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b74978708bb95475562412d4654c00fbcedd9f89/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/ImpactedResources_Get.json
func ExampleImpactedResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewImpactedResourcesClient().Get(ctx, "BC_1-FXZ", "abc-123-ghj-456", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EventImpactedResource = armresourcehealth.EventImpactedResource{
	// 	Name: to.Ptr("abc-123-ghj-456"),
	// 	Type: to.Ptr("Microsoft.ResourceHealth/events/impactedResources"),
	// 	ID: to.Ptr("/subscriptions/{subscripitionId}/providers/Microsoft.ResourceHealth/events/BC_1-FXZ/impactedResources/abc-123-ghj-456"),
	// 	Properties: &armresourcehealth.EventImpactedResourceProperties{
	// 		TargetRegion: to.Ptr("westus"),
	// 		TargetResourceID: to.Ptr("/subscriptions/4970d23e-ed41-4670-9c19-02a1d2808dd9/resourceGroups/TEST/providers/Microsoft.Compute/virtualMachines/testvm"),
	// 		TargetResourceType: to.Ptr("Microsoft.Compute/VirtualMachines"),
	// 	},
	// }
}
