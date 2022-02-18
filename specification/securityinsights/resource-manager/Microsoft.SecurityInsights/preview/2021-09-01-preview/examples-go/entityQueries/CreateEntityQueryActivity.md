Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsight%2Farmsecurityinsight%2Fv0.2.1/sdk/resourcemanager/securityinsight/armsecurityinsight/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/entityQueries/CreateEntityQueryActivity.json
func ExampleEntityQueriesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewEntityQueriesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<entity-query-id>",
		&armsecurityinsight.ActivityCustomEntityQuery{
			Etag: to.StringPtr("<etag>"),
			Kind: armsecurityinsight.CustomEntityQueryKind("Activity").ToPtr(),
			Properties: &armsecurityinsight.ActivityEntityQueriesProperties{
				Description: to.StringPtr("<description>"),
				Content:     to.StringPtr("<content>"),
				Enabled:     to.BoolPtr(true),
				EntitiesFilter: map[string][]*string{
					"Host_OsFamily": {
						to.StringPtr("Windows")},
				},
				InputEntityType: armsecurityinsight.EntityType("Host").ToPtr(),
				QueryDefinitions: &armsecurityinsight.ActivityEntityQueriesPropertiesQueryDefinitions{
					Query: to.StringPtr("<query>"),
				},
				RequiredInputFieldsSets: [][]*string{
					{
						to.StringPtr("Host_HostName"),
						to.StringPtr("Host_NTDomain")},
					{
						to.StringPtr("Host_HostName"),
						to.StringPtr("Host_DnsDomain")},
					{
						to.StringPtr("Host_AzureID")},
					{
						to.StringPtr("Host_OMSAgentID")}},
				Title: to.StringPtr("<title>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EntityQueriesClientCreateOrUpdateResult)
}
```
