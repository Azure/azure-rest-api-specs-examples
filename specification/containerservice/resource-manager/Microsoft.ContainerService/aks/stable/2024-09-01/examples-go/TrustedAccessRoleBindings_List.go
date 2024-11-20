package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e838027e88cca634c1545e744630de9262a6e72a/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2024-09-01/examples/TrustedAccessRoleBindings_List.json
func ExampleTrustedAccessRoleBindingsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTrustedAccessRoleBindingsClient().NewListPager("rg1", "clustername1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.TrustedAccessRoleBindingListResult = armcontainerservice.TrustedAccessRoleBindingListResult{
		// 	Value: []*armcontainerservice.TrustedAccessRoleBinding{
		// 		{
		// 			Name: to.Ptr("binding1"),
		// 			Type: to.Ptr("Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/trustedAccessRoleBindings/binding1"),
		// 			Properties: &armcontainerservice.TrustedAccessRoleBindingProperties{
		// 				Roles: []*string{
		// 					to.Ptr("Microsoft.MachineLearningServices/workspaces/reader"),
		// 					to.Ptr("Microsoft.MachineLearningServices/workspaces/writer")},
		// 					SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/b/providers/Microsoft.MachineLearningServices/workspaces/c"),
		// 				},
		// 		}},
		// 	}
	}
}
