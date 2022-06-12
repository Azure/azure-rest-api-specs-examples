```go
package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/HyperVCollectors_Create.json
func ExampleHyperVCollectorsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmigrate.NewHyperVCollectorsClient("8c3c936a-c09b-4de3-830b-3f5f244d72e9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"contosoithyperv",
		"migrateprojectce73project",
		"migrateprojectce73collector",
		&armmigrate.HyperVCollectorsClientCreateOptions{CollectorBody: &armmigrate.HyperVCollector{
			ETag: to.Ptr("\"00000981-0000-0300-0000-5d74cd5f0000\""),
			Properties: &armmigrate.CollectorProperties{
				AgentProperties: &armmigrate.CollectorAgentProperties{
					SpnDetails: &armmigrate.CollectorBodyAgentSpnProperties{
						ApplicationID: to.Ptr("827f1053-44dc-439f-b832-05416dcce12b"),
						Audience:      to.Ptr("https://72f988bf-86f1-41af-91ab-2d7cd011db47/migrateprojectce73agentauthaadapp"),
						Authority:     to.Ptr("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
						ObjectID:      to.Ptr("be75098e-c0fc-4ac4-98c7-282ebbcf8370"),
						TenantID:      to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
					},
				},
				DiscoverySiteID: to.Ptr("/subscriptions/8c3c936a-c09b-4de3-830b-3f5f244d72e9/resourceGroups/ContosoITHyperV/providers/Microsoft.OffAzure/HyperVSites/migrateprojectce73site"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmigrate%2Farmmigrate%2Fv1.0.0/sdk/resourcemanager/migrate/armmigrate/README.md) on how to add the SDK to your project and authenticate.
