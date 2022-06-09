```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobTargetGroupMax.json
func ExampleJobTargetGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewJobTargetGroupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"group1",
		"server1",
		"agent1",
		"targetGroup1",
		armsql.JobTargetGroup{
			Properties: &armsql.JobTargetGroupProperties{
				Members: []*armsql.JobTarget{
					{
						Type:           to.Ptr(armsql.JobTargetTypeSQLDatabase),
						DatabaseName:   to.Ptr("database1"),
						MembershipType: to.Ptr(armsql.JobTargetGroupMembershipTypeExclude),
						ServerName:     to.Ptr("server1"),
					},
					{
						Type:              to.Ptr(armsql.JobTargetTypeSQLServer),
						MembershipType:    to.Ptr(armsql.JobTargetGroupMembershipTypeInclude),
						RefreshCredential: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/testCredential"),
						ServerName:        to.Ptr("server1"),
					},
					{
						Type:              to.Ptr(armsql.JobTargetTypeSQLElasticPool),
						ElasticPoolName:   to.Ptr("pool1"),
						MembershipType:    to.Ptr(armsql.JobTargetGroupMembershipTypeInclude),
						RefreshCredential: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/testCredential"),
						ServerName:        to.Ptr("server2"),
					},
					{
						Type:              to.Ptr(armsql.JobTargetTypeSQLShardMap),
						MembershipType:    to.Ptr(armsql.JobTargetGroupMembershipTypeInclude),
						RefreshCredential: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/testCredential"),
						ServerName:        to.Ptr("server3"),
						ShardMapName:      to.Ptr("shardMap1"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv1.0.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.
