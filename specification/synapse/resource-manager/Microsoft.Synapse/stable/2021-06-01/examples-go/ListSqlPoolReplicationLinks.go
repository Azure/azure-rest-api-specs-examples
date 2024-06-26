package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolReplicationLinks.json
func ExampleSQLPoolReplicationLinksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLPoolReplicationLinksClient().NewListPager("sqlcrudtest-4799", "sqlcrudtest-6440", "testdb", nil)
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
		// page.ReplicationLinkListResult = armsynapse.ReplicationLinkListResult{
		// 	Value: []*armsynapse.ReplicationLink{
		// 		{
		// 			Name: to.Ptr("5b301b68-03f6-4b26-b0f4-73ebb8634238"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/replicationLinks"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-4799/providers/Microsoft.Synapse/workspaces/sqlcrudtest-6440/sqlPools/testdb/replicationLinks/5b301b68-03f6-4b26-b0f4-73ebb8634238"),
		// 			Location: to.Ptr("Japan East"),
		// 			Properties: &armsynapse.ReplicationLinkProperties{
		// 				IsTerminationAllowed: to.Ptr(true),
		// 				PartnerDatabase: to.Ptr("testdb"),
		// 				PartnerLocation: to.Ptr("Japan East"),
		// 				PartnerRole: to.Ptr(armsynapse.ReplicationRolePrimary),
		// 				PartnerServer: to.Ptr("sqlcrudtest-5961"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				ReplicationMode: to.Ptr("ASYNC"),
		// 				ReplicationState: to.Ptr(armsynapse.ReplicationStateCATCHUP),
		// 				Role: to.Ptr(armsynapse.ReplicationRoleSecondary),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-10T01:44:27.117Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
