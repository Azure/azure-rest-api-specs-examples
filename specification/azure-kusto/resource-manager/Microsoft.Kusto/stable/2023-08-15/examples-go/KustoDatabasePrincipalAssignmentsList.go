package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabasePrincipalAssignmentsList.json
func ExampleDatabasePrincipalAssignmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabasePrincipalAssignmentsClient().NewListPager("kustorptest", "kustoCluster", "Kustodatabase8", nil)
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
		// page.DatabasePrincipalAssignmentListResult = armkusto.DatabasePrincipalAssignmentListResult{
		// 	Value: []*armkusto.DatabasePrincipalAssignment{
		// 		{
		// 			Name: to.Ptr("kustoCluster/Kustodatabase8/kustoprincipal1"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/PrincipalAssignments"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/Kustodatabase8/PrincipalAssignments/kustoprincipal1"),
		// 			Properties: &armkusto.DatabasePrincipalProperties{
		// 				AADObjectID: to.Ptr("98765432-1234-1234-1234-123456789123"),
		// 				PrincipalID: to.Ptr("11223344-1234-1234-1234-123456789999"),
		// 				PrincipalName: to.Ptr("TestApp"),
		// 				PrincipalType: to.Ptr(armkusto.PrincipalTypeApp),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 				Role: to.Ptr(armkusto.DatabasePrincipalRoleViewer),
		// 				TenantID: to.Ptr("12345678-1234-1234-1234-123456789123"),
		// 				TenantName: to.Ptr("tenantName"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kustoCluster/Kustodatabase8/kustoprincipal2"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/PrincipalAssignments"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/Kustodatabase8/PrincipalAssignments/kustoprincipal1"),
		// 			Properties: &armkusto.DatabasePrincipalProperties{
		// 				AADObjectID: to.Ptr("98765432-1234-1234-1234-123456789123"),
		// 				PrincipalID: to.Ptr("87654321-1234-1234-1234-123456789123"),
		// 				PrincipalName: to.Ptr("TestApp"),
		// 				PrincipalType: to.Ptr(armkusto.PrincipalTypeApp),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 				Role: to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
		// 				TenantID: to.Ptr("12345678-1234-1234-1234-123456789123"),
		// 				TenantName: to.Ptr("tenantName"),
		// 			},
		// 	}},
		// }
	}
}
