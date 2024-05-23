package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/FailoverGroupList.json
func ExampleFailoverGroupsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFailoverGroupsClient().NewListByServerPager("Default", "failover-group-primary-server", nil)
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
		// page.FailoverGroupListResult = armsql.FailoverGroupListResult{
		// 	Value: []*armsql.FailoverGroup{
		// 		{
		// 			Name: to.Ptr("failover-group-test"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/failoverGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/failoverGroups/failover-group-test"),
		// 			Location: to.Ptr("Japan East"),
		// 			Properties: &armsql.FailoverGroupProperties{
		// 				Databases: []*string{
		// 				},
		// 				PartnerServers: []*armsql.PartnerInfo{
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"),
		// 						Location: to.Ptr("Japan West"),
		// 						ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRoleSecondary),
		// 				}},
		// 				ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
		// 					FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
		// 					TargetServer: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"),
		// 				},
		// 				ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
		// 					FailoverPolicy: to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
		// 					FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](480),
		// 				},
		// 				ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRolePrimary),
		// 				ReplicationState: to.Ptr("CATCH_UP"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("failover-group-test-2"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/failoverGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/failoverGroups/failover-group-test-2"),
		// 			Location: to.Ptr("Japan East"),
		// 			Properties: &armsql.FailoverGroupProperties{
		// 				Databases: []*string{
		// 				},
		// 				PartnerServers: []*armsql.PartnerInfo{
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"),
		// 						Location: to.Ptr("Japan West"),
		// 						ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRoleSecondary),
		// 				}},
		// 				ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
		// 					FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
		// 					TargetServer: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"),
		// 				},
		// 				ReadWriteEndpoint: &armsql.FailoverGroupReadWriteEndpoint{
		// 					FailoverPolicy: to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
		// 					FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](480),
		// 				},
		// 				ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRolePrimary),
		// 				ReplicationState: to.Ptr("CATCH_UP"),
		// 			},
		// 	}},
		// }
	}
}
