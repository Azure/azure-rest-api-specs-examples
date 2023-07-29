package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/InstanceFailoverGroupList.json
func ExampleInstanceFailoverGroupsClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInstanceFailoverGroupsClient().NewListByLocationPager("Default", "Japan East", nil)
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
		// page.InstanceFailoverGroupListResult = armsql.InstanceFailoverGroupListResult{
		// 	Value: []*armsql.InstanceFailoverGroup{
		// 		{
		// 			Name: to.Ptr("failover-group-test"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/instanceFailoverGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/locations/JapanEast/instanceFailoverGroups/failover-group-test"),
		// 			Properties: &armsql.InstanceFailoverGroupProperties{
		// 				ManagedInstancePairs: []*armsql.ManagedInstancePairInfo{
		// 					{
		// 						PartnerManagedInstanceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/failover-group-secondary-mngdInstance"),
		// 						PrimaryManagedInstanceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/failover-group-primary-mngdInstance"),
		// 				}},
		// 				PartnerRegions: []*armsql.PartnerRegionInfo{
		// 					{
		// 						Location: to.Ptr("Japan West"),
		// 						ReplicationRole: to.Ptr(armsql.InstanceFailoverGroupReplicationRoleSecondary),
		// 				}},
		// 				ReadOnlyEndpoint: &armsql.InstanceFailoverGroupReadOnlyEndpoint{
		// 					FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
		// 				},
		// 				ReadWriteEndpoint: &armsql.InstanceFailoverGroupReadWriteEndpoint{
		// 					FailoverPolicy: to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
		// 					FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](480),
		// 				},
		// 				ReplicationRole: to.Ptr(armsql.InstanceFailoverGroupReplicationRolePrimary),
		// 				ReplicationState: to.Ptr("CATCH_UP"),
		// 				SecondaryType: to.Ptr(armsql.SecondaryInstanceTypeGeo),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("failover-group-test-1"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/instanceFailoverGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/locations/JapanEast/instanceFailoverGroups/failover-group-test-1"),
		// 			Properties: &armsql.InstanceFailoverGroupProperties{
		// 				ManagedInstancePairs: []*armsql.ManagedInstancePairInfo{
		// 					{
		// 						PartnerManagedInstanceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/failover-group-secondary-mngdInstance-1"),
		// 						PrimaryManagedInstanceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/failover-group-primary-mngdInstance-1"),
		// 				}},
		// 				PartnerRegions: []*armsql.PartnerRegionInfo{
		// 					{
		// 						Location: to.Ptr("Japan West"),
		// 						ReplicationRole: to.Ptr(armsql.InstanceFailoverGroupReplicationRoleSecondary),
		// 				}},
		// 				ReadOnlyEndpoint: &armsql.InstanceFailoverGroupReadOnlyEndpoint{
		// 					FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyDisabled),
		// 				},
		// 				ReadWriteEndpoint: &armsql.InstanceFailoverGroupReadWriteEndpoint{
		// 					FailoverPolicy: to.Ptr(armsql.ReadWriteEndpointFailoverPolicyAutomatic),
		// 					FailoverWithDataLossGracePeriodMinutes: to.Ptr[int32](480),
		// 				},
		// 				ReplicationRole: to.Ptr(armsql.InstanceFailoverGroupReplicationRolePrimary),
		// 				ReplicationState: to.Ptr("CATCH_UP"),
		// 				SecondaryType: to.Ptr(armsql.SecondaryInstanceTypeGeo),
		// 			},
		// 	}},
		// }
	}
}
