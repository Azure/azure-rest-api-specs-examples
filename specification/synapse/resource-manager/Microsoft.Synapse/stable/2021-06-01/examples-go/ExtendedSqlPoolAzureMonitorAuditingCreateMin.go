package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ExtendedSqlPoolAzureMonitorAuditingCreateMin.json
func ExampleExtendedSQLPoolBlobAuditingPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewExtendedSQLPoolBlobAuditingPoliciesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"blobauditingtest-4799",
		"blobauditingtest-6440",
		"testdb",
		armsynapse.ExtendedSQLPoolBlobAuditingPolicy{
			Properties: &armsynapse.ExtendedSQLPoolBlobAuditingPolicyProperties{
				IsAzureMonitorTargetEnabled: to.Ptr(true),
				State:                       to.Ptr(armsynapse.BlobAuditingPolicyStateEnabled),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
