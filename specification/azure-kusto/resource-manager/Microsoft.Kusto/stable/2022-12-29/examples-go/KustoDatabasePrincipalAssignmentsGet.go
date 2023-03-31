package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoDatabasePrincipalAssignmentsGet.json
func ExampleDatabasePrincipalAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabasePrincipalAssignmentsClient().Get(ctx, "kustorptest", "kustoCluster", "Kustodatabase8", "kustoprincipal1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
