package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8358c7473dfe057d84a6b6a921225063c040b31a/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/FailoverGroupGet.json
func ExampleFailoverGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFailoverGroupsClient().Get(ctx, "Default", "failovergroupprimaryserver", "failovergrouptest3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FailoverGroup = armsql.FailoverGroup{
	// 	Name: to.Ptr("failovergrouptest3"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/failoverGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failovergroupprimaryserver/failoverGroups/failovergrouptest3"),
	// 	Location: to.Ptr("Japan West"),
	// 	Properties: &armsql.FailoverGroupProperties{
	// 		Databases: []*string{
	// 			to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failovergroupprimaryserver/databases/testdb1"),
	// 			to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failovergroupprimaryserver/databases/testdb2")},
	// 			PartnerServers: []*armsql.PartnerInfo{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failovergroupsecondaryserver"),
	// 					Location: to.Ptr("Japan East"),
	// 					ReplicationRole: to.Ptr(armsql.FailoverGroupReplicationRoleSecondary),
	// 			}},
	// 			ReadOnlyEndpoint: &armsql.FailoverGroupReadOnlyEndpoint{
	// 				FailoverPolicy: to.Ptr(armsql.ReadOnlyEndpointFailoverPolicyEnabled),
	// 				TargetServer: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failovergroupsecondaryserver"),
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
