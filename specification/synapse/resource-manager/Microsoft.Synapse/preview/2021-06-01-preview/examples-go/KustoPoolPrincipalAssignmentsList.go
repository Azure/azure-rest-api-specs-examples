package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolPrincipalAssignmentsList.json
func ExampleKustoPoolPrincipalAssignmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKustoPoolPrincipalAssignmentsClient().NewListPager("synapseWorkspaceName", "kustoclusterrptest4", "kustorptest", nil)
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
		// page.ClusterPrincipalAssignmentListResult = armsynapse.ClusterPrincipalAssignmentListResult{
		// 	Value: []*armsynapse.ClusterPrincipalAssignment{
		// 		{
		// 			Name: to.Ptr("synapseWorkspaceName/kustoclusterrptest4/kustoprincipal1"),
		// 			Type: to.Ptr("Microsoft.Synapse/Workspaces/KustoPools/PrincipalAssignments"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/synapseWorkspaceName/kustoPools/kustoclusterrptest4/PrincipalAssignments/kustoprincipal1"),
		// 			Properties: &armsynapse.ClusterPrincipalProperties{
		// 				AADObjectID: to.Ptr("98765432-1234-1234-1234-123456789123"),
		// 				PrincipalID: to.Ptr("87654321-1234-1234-1234-123456789123"),
		// 				PrincipalName: to.Ptr("TestApp"),
		// 				PrincipalType: to.Ptr(armsynapse.PrincipalTypeApp),
		// 				ProvisioningState: to.Ptr(armsynapse.ResourceProvisioningStateSucceeded),
		// 				Role: to.Ptr(armsynapse.ClusterPrincipalRole("Admin")),
		// 				TenantID: to.Ptr("12345678-1234-1234-1234-123456789123"),
		// 				TenantName: to.Ptr("tenantName"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("synapseWorkspaceName/kustoclusterrptest4/kustoprincipal2"),
		// 			Type: to.Ptr("Microsoft.Synapse/Workspaces/KustoPools/PrincipalAssignments"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/synapseWorkspaceName/kustoPools/kustoclusterrptest4/PrincipalAssignments/kustoprincipal2"),
		// 			Properties: &armsynapse.ClusterPrincipalProperties{
		// 				AADObjectID: to.Ptr("98765432-1234-1234-1234-123456789123"),
		// 				PrincipalID: to.Ptr("87654321-1234-1234-1234-123456789123"),
		// 				PrincipalName: to.Ptr("TestApp"),
		// 				PrincipalType: to.Ptr(armsynapse.PrincipalTypeApp),
		// 				ProvisioningState: to.Ptr(armsynapse.ResourceProvisioningStateSucceeded),
		// 				Role: to.Ptr(armsynapse.ClusterPrincipalRole("Admin")),
		// 				TenantID: to.Ptr("12345678-1234-1234-1234-123456789123"),
		// 				TenantName: to.Ptr("tenantName"),
		// 			},
		// 	}},
		// }
	}
}
