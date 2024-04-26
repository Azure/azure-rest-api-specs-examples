package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1dd99306d14fd6c602f47652a209a4a6812c368c/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-02-02-preview/examples/TrustedAccessRoleBindings_Get.json
func ExampleTrustedAccessRoleBindingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTrustedAccessRoleBindingsClient().Get(ctx, "rg1", "clustername1", "binding1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TrustedAccessRoleBinding = armcontainerservice.TrustedAccessRoleBinding{
	// 	Name: to.Ptr("binding1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/trustedAccessRoleBindings/binding1"),
	// 	Properties: &armcontainerservice.TrustedAccessRoleBindingProperties{
	// 		Roles: []*string{
	// 			to.Ptr("Microsoft.MachineLearningServices/workspaces/reader"),
	// 			to.Ptr("Microsoft.MachineLearningServices/workspaces/writer")},
	// 			SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/b/providers/Microsoft.MachineLearningServices/workspaces/c"),
	// 		},
	// 	}
}
