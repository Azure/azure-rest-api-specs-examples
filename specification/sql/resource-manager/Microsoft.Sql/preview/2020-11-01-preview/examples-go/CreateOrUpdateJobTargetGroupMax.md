Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.5.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

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
		return
	}
	ctx := context.Background()
	client, err := armsql.NewJobTargetGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<job-agent-name>",
		"<target-group-name>",
		armsql.JobTargetGroup{
			Properties: &armsql.JobTargetGroupProperties{
				Members: []*armsql.JobTarget{
					{
						Type:           to.Ptr(armsql.JobTargetTypeSQLDatabase),
						DatabaseName:   to.Ptr("<database-name>"),
						MembershipType: to.Ptr(armsql.JobTargetGroupMembershipTypeExclude),
						ServerName:     to.Ptr("<server-name>"),
					},
					{
						Type:              to.Ptr(armsql.JobTargetTypeSQLServer),
						MembershipType:    to.Ptr(armsql.JobTargetGroupMembershipTypeInclude),
						RefreshCredential: to.Ptr("<refresh-credential>"),
						ServerName:        to.Ptr("<server-name>"),
					},
					{
						Type:              to.Ptr(armsql.JobTargetTypeSQLElasticPool),
						ElasticPoolName:   to.Ptr("<elastic-pool-name>"),
						MembershipType:    to.Ptr(armsql.JobTargetGroupMembershipTypeInclude),
						RefreshCredential: to.Ptr("<refresh-credential>"),
						ServerName:        to.Ptr("<server-name>"),
					},
					{
						Type:              to.Ptr(armsql.JobTargetTypeSQLShardMap),
						MembershipType:    to.Ptr(armsql.JobTargetGroupMembershipTypeInclude),
						RefreshCredential: to.Ptr("<refresh-credential>"),
						ServerName:        to.Ptr("<server-name>"),
						ShardMapName:      to.Ptr("<shard-map-name>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
