package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/ServerCollectors_Create.json
func ExampleServerCollectorsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmigrate.NewServerCollectorsClient("4bd2aa0f-2bd2-4d67-91a8-5a4533d58600", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"pajindtest",
		"app11141project",
		"app23df4collector",
		&armmigrate.ServerCollectorsClientCreateOptions{CollectorBody: &armmigrate.ServerCollector{
			ETag: to.Ptr("\"00000606-0000-0d00-0000-605999bf0000\""),
			Properties: &armmigrate.CollectorProperties{
				AgentProperties: &armmigrate.CollectorAgentProperties{
					SpnDetails: &armmigrate.CollectorBodyAgentSpnProperties{
						ApplicationID: to.Ptr("ad9f701a-cc08-4421-b51f-b5762d58e9ba"),
						Audience:      to.Ptr("https://72f988bf-86f1-41af-91ab-2d7cd011db47/app23df4authandaccessaadapp"),
						Authority:     to.Ptr("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
						ObjectID:      to.Ptr("b4975e42-9248-4a36-b99f-37eca377ea00"),
						TenantID:      to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
					},
				},
				DiscoverySiteID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/pajindTest/providers/Microsoft.OffAzure/ServerSites/app21141site"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
