```go
package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/entityQueries/CreateEntityQueryActivity.json
func ExampleEntityQueriesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewEntityQueriesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<entity-query-id>",
		&armsecurityinsights.ActivityCustomEntityQuery{
			Etag: to.Ptr("<etag>"),
			Kind: to.Ptr(armsecurityinsights.CustomEntityQueryKindActivity),
			Properties: &armsecurityinsights.ActivityEntityQueriesProperties{
				Description: to.Ptr("<description>"),
				Content:     to.Ptr("<content>"),
				Enabled:     to.Ptr(true),
				EntitiesFilter: map[string][]*string{
					"Host_OsFamily": {
						to.Ptr("Windows")},
				},
				InputEntityType: to.Ptr(armsecurityinsights.EntityTypeHost),
				QueryDefinitions: &armsecurityinsights.ActivityEntityQueriesPropertiesQueryDefinitions{
					Query: to.Ptr("<query>"),
				},
				RequiredInputFieldsSets: [][]*string{
					{
						to.Ptr("Host_HostName"),
						to.Ptr("Host_NTDomain")},
					{
						to.Ptr("Host_HostName"),
						to.Ptr("Host_DnsDomain")},
					{
						to.Ptr("Host_AzureID")},
					{
						to.Ptr("Host_OMSAgentID")}},
				Title: to.Ptr("<title>"),
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsights%2Farmsecurityinsights%2Fv0.3.0/sdk/resourcemanager/securityinsights/armsecurityinsights/README.md) on how to add the SDK to your project and authenticate.
