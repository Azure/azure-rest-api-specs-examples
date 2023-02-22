package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasePrincipalAssignmentsCreateOrUpdate.json
func ExampleKustoPoolDatabasePrincipalAssignmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolDatabasePrincipalAssignmentsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "synapseWorkspaceName", "kustoclusterrptest4", "Kustodatabase8", "kustoprincipal1", "kustorptest", armsynapse.DatabasePrincipalAssignment{
		Properties: &armsynapse.DatabasePrincipalProperties{
			PrincipalID:   to.Ptr("87654321-1234-1234-1234-123456789123"),
			PrincipalType: to.Ptr(armsynapse.PrincipalTypeApp),
			Role:          to.Ptr(armsynapse.DatabasePrincipalRoleAdmin),
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
	// res.DatabasePrincipalAssignment = armsynapse.DatabasePrincipalAssignment{
	// 	Name: to.Ptr("synapseWorkspaceName/kustoclusterrptest4/Kustodatabase8/kustoprincipal1"),
	// 	Type: to.Ptr("Microsoft.Synapse/Workspaces/KustoPools/Databases/PrincipalAssignments"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/synapseWorkspaceName/kustoPools/kustoclusterrptest4/Databases/Kustodatabase8/PrincipalAssignments/kustoprincipal1"),
	// 	Properties: &armsynapse.DatabasePrincipalProperties{
	// 		AADObjectID: to.Ptr("98765432-1234-1234-1234-123456789123"),
	// 		PrincipalID: to.Ptr("87654321-1234-1234-1234-123456789123"),
	// 		PrincipalName: to.Ptr("TestApp"),
	// 		PrincipalType: to.Ptr(armsynapse.PrincipalTypeApp),
	// 		ProvisioningState: to.Ptr(armsynapse.ResourceProvisioningStateSucceeded),
	// 		Role: to.Ptr(armsynapse.DatabasePrincipalRoleAdmin),
	// 		TenantID: to.Ptr("12345678-1234-1234-1234-123456789123"),
	// 		TenantName: to.Ptr("tenantName"),
	// 	},
	// }
}
