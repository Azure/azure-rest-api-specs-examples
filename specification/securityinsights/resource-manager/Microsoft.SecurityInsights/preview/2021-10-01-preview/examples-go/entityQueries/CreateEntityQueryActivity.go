package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/entityQueries/CreateEntityQueryActivity.json
func ExampleEntityQueriesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewEntityQueriesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<entity-query-id>",
		&armsecurityinsights.ActivityCustomEntityQuery{
			Etag: to.StringPtr("<etag>"),
			Kind: armsecurityinsights.CustomEntityQueryKind("Activity").ToPtr(),
			Properties: &armsecurityinsights.ActivityEntityQueriesProperties{
				Description: to.StringPtr("<description>"),
				Content:     to.StringPtr("<content>"),
				Enabled:     to.BoolPtr(true),
				EntitiesFilter: map[string][]*string{
					"Host_OsFamily": {
						to.StringPtr("Windows")},
				},
				InputEntityType: armsecurityinsights.EntityType("Host").ToPtr(),
				QueryDefinitions: &armsecurityinsights.ActivityEntityQueriesPropertiesQueryDefinitions{
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
