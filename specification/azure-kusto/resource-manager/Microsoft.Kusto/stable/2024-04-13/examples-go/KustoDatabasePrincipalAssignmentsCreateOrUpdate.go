package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8adce17dc500f338f86f18af30aac61dcb71c5f/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoDatabasePrincipalAssignmentsCreateOrUpdate.json
func ExampleDatabasePrincipalAssignmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasePrincipalAssignmentsClient().BeginCreateOrUpdate(ctx, "kustorptest", "kustoCluster", "Kustodatabase8", "kustoprincipal1", armkusto.DatabasePrincipalAssignment{
		Properties: &armkusto.DatabasePrincipalProperties{
			PrincipalID:   to.Ptr("87654321-1234-1234-1234-123456789123"),
			PrincipalType: to.Ptr(armkusto.PrincipalTypeApp),
			Role:          to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
			TenantID:      to.Ptr("12345678-1234-1234-1234-123456789123"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabasePrincipalAssignment = armkusto.DatabasePrincipalAssignment{
	// 	Name: to.Ptr("kustoCluster/Kustodatabase8/kustoprincipal1"),
	// 	Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/PrincipalAssignments"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/Kustodatabase8/PrincipalAssignments/kustoprincipal1"),
	// 	Properties: &armkusto.DatabasePrincipalProperties{
	// 		AADObjectID: to.Ptr("98765432-1234-1234-1234-123456789123"),
	// 		PrincipalID: to.Ptr("87654321-1234-1234-1234-123456789123"),
	// 		PrincipalName: to.Ptr("TestApp"),
	// 		PrincipalType: to.Ptr(armkusto.PrincipalTypeApp),
	// 		ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
	// 		Role: to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
	// 		TenantID: to.Ptr("12345678-1234-1234-1234-123456789123"),
	// 		TenantName: to.Ptr("tenantName"),
	// 	},
	// }
}
