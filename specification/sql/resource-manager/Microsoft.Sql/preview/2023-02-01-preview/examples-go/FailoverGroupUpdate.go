package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/FailoverGroupUpdate.json
func ExampleFailoverGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFailoverGroupsClient().BeginUpdate(ctx, "Default", "failover-group-primary-server", "failover-group-test-1", armsql.FailoverGroupUpdate{
		Properties: &armsql.FailoverGroupUpdateProperties{
			Databases: []*string{
				to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-1")},
			ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
				FailoverPolicy:                         to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
				FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](120),
			},
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
	// res.FailoverGroup = armsql.FailoverGroup{
	// 	Name: to.Ptr("failover-group-test-3"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/failoverGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/failoverGroups/failover-group-test-3"),
	// 	Location: to.Ptr("Japan East"),
	// 	Properties: &armsql.FailoverGroupProperties{
	// 		Databases: []*string{
	// 			to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-1")},
	// 			PartnerServers: []*armsql.PartnerInfo{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"),
	// 					Location: to.Ptr("Japan West"),
	// 					ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRoleSecondary),
	// 			}},
	// 			ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
	// 				FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
	// 				TargetServer: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"),
	// 			},
	// 			ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
	// 				FailoverPolicy: to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
	// 				FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](120),
	// 			},
	// 			ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRolePrimary),
	// 			ReplicationState: to.Ptr("CATCH_UP"),
	// 		},
	// 	}
}
