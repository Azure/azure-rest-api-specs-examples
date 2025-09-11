package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8358c7473dfe057d84a6b6a921225063c040b31a/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/ReplicationLinkFailover.json
func ExampleReplicationLinksClient_BeginFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationLinksClient().BeginFailover(ctx, "Default", "sourcesvr", "gamma-db", "4891ca10-ebd0-47d7-9182-c722651780fb", nil)
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
	// res.ReplicationLink = armsql.ReplicationLink{
	// 	Name: to.Ptr("4891ca10-ebd0-47d7-9182-c722651780fb"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sourcesvr/databases/gamma-db/replicationLinks/4891ca10-ebd0-47d7-9182-c722651780fb"),
	// 	Properties: &armsql.ReplicationLinkProperties{
	// 		IsTerminationAllowed: to.Ptr(true),
	// 		LinkType: to.Ptr(armsql.ReplicationLinkTypeGEO),
	// 		PartnerDatabase: to.Ptr("gamma-db"),
	// 		PartnerDatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-555555555555/resourceGroups/Second-Default/providers/Microsoft.Sql/servers/testsvr/databases/gamma-db"),
	// 		PartnerLocation: to.Ptr("Japan East"),
	// 		PartnerRole: to.Ptr(armsql.ReplicationRoleSecondary),
	// 		PartnerServer: to.Ptr("testsvr"),
	// 		PercentComplete: to.Ptr[int32](100),
	// 		ReplicationMode: to.Ptr("ASYNC"),
	// 		ReplicationState: to.Ptr(armsql.ReplicationStateCATCHUP),
	// 		Role: to.Ptr(armsql.ReplicationRolePrimary),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-21T08:12:43.783Z"); return t}()),
	// 	},
	// }
}
