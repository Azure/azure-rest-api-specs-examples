Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.3.1/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobTargetGroupMax.json
func ExampleJobTargetGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewJobTargetGroupsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<job-agent-name>",
		"<target-group-name>",
		armsql.JobTargetGroup{
			Properties: &armsql.JobTargetGroupProperties{
				Members: []*armsql.JobTarget{
					{
						Type:           armsql.JobTargetType("SqlDatabase").ToPtr(),
						DatabaseName:   to.StringPtr("<database-name>"),
						MembershipType: armsql.JobTargetGroupMembershipTypeExclude.ToPtr(),
						ServerName:     to.StringPtr("<server-name>"),
					},
					{
						Type:              armsql.JobTargetType("SqlServer").ToPtr(),
						MembershipType:    armsql.JobTargetGroupMembershipTypeInclude.ToPtr(),
						RefreshCredential: to.StringPtr("<refresh-credential>"),
						ServerName:        to.StringPtr("<server-name>"),
					},
					{
						Type:              armsql.JobTargetType("SqlElasticPool").ToPtr(),
						ElasticPoolName:   to.StringPtr("<elastic-pool-name>"),
						MembershipType:    armsql.JobTargetGroupMembershipTypeInclude.ToPtr(),
						RefreshCredential: to.StringPtr("<refresh-credential>"),
						ServerName:        to.StringPtr("<server-name>"),
					},
					{
						Type:              armsql.JobTargetType("SqlShardMap").ToPtr(),
						MembershipType:    armsql.JobTargetGroupMembershipTypeInclude.ToPtr(),
						RefreshCredential: to.StringPtr("<refresh-credential>"),
						ServerName:        to.StringPtr("<server-name>"),
						ShardMapName:      to.StringPtr("<shard-map-name>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.JobTargetGroupsClientCreateOrUpdateResult)
}
```
