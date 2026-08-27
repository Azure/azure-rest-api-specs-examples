package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v3"
)

// Generated from example definition: 2026-08-01-preview/DiscoveredResources_Get.json
func ExampleDiscoveredResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiscoveredResourcesClient().Get(ctx, "exampleRG", "exampleWorkspace", "a1b2c3d4-e5f6-7890-abcd-ef1234567890", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.DiscoveredResourcesClientGetResponse{
	// 	DiscoveredResource: armchaos.DiscoveredResource{
	// 		ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/discoveredResources/a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
	// 		Name: to.Ptr("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
	// 		Type: to.Ptr("Microsoft.Chaos/workspaces/discoveredResources"),
	// 		Properties: &armchaos.DiscoveredResourceProperties{
	// 			ResourceNamespace: to.Ptr("Microsoft.Compute"),
	// 			ResourceName: to.Ptr("myVirtualMachine"),
	// 			ResourceType: to.Ptr("virtualMachines"),
	// 			FullyQualifiedIdentifier: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/myVirtualMachine"),
	// 			DiscoveredAt: to.Ptr(time.Date(2025, time.January, 15, 10, 30, 0, 0, time.UTC)),
	// 			Scope: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG"),
	// 		},
	// 		SystemData: &armchaos.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2025, time.January, 15, 10, 30, 0, 0, time.UTC)),
	// 			CreatedBy: to.Ptr("system"),
	// 			CreatedByType: to.Ptr(armchaos.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(time.Date(2025, time.January, 15, 10, 30, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("system"),
	// 			LastModifiedByType: to.Ptr(armchaos.CreatedByTypeApplication),
	// 		},
	// 	},
	// }
}
