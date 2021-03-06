package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolPrincipalAssignmentsCreateOrUpdate.json
func ExampleKustoPoolPrincipalAssignmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolPrincipalAssignmentsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"synapseWorkspaceName",
		"kustoclusterrptest4",
		"kustoprincipal1",
		"kustorptest",
		armsynapse.ClusterPrincipalAssignment{
			Properties: &armsynapse.ClusterPrincipalProperties{
				PrincipalID:   to.Ptr("87654321-1234-1234-1234-123456789123"),
				PrincipalType: to.Ptr(armsynapse.PrincipalTypeApp),
				Role:          to.Ptr(armsynapse.ClusterPrincipalRoleAllDatabasesAdmin),
				TenantID:      to.Ptr("12345678-1234-1234-1234-123456789123"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
