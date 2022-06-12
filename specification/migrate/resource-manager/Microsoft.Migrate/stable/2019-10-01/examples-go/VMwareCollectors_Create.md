```go
package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/VMwareCollectors_Create.json
func ExampleVMwareCollectorsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmigrate.NewVMwareCollectorsClient("6393a73f-8d55-47ef-b6dd-179b3e0c7910", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"abgoyal-westEurope",
		"abgoyalWEselfhostb72bproject",
		"PortalvCenterbc2fcollector",
		&armmigrate.VMwareCollectorsClientCreateOptions{CollectorBody: &armmigrate.VMwareCollector{
			ETag: to.Ptr("\"01003d32-0000-0d00-0000-5d74d2e50000\""),
			Properties: &armmigrate.CollectorProperties{
				AgentProperties: &armmigrate.CollectorAgentProperties{
					SpnDetails: &armmigrate.CollectorBodyAgentSpnProperties{
						ApplicationID: to.Ptr("fc717575-8173-4b21-92a5-658b655e613e"),
						Audience:      to.Ptr("https://72f988bf-86f1-41af-91ab-2d7cd011db47/PortalvCenterbc2fagentauthaadapp"),
						Authority:     to.Ptr("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
						ObjectID:      to.Ptr("29d94f38-db94-4980-aec0-0cfd55ab1cd0"),
						TenantID:      to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
					},
				},
				DiscoverySiteID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westEurope/providers/Microsoft.OffAzure/VMwareSites/PortalvCenterbc2fsite"),
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
