package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/HyperVCollectors_Create.json
func ExampleHyperVCollectorsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHyperVCollectorsClient().Create(ctx, "contosoithyperv", "migrateprojectce73project", "migrateprojectce73collector", &armmigrate.HyperVCollectorsClientCreateOptions{CollectorBody: &armmigrate.HyperVCollector{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.HyperVCollector = armmigrate.HyperVCollector{
	// 	Name: to.Ptr("migrateprojectce73collector"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/hypervcollectors"),
	// 	ETag: to.Ptr("\"00000981-0000-0300-0000-5d74cd5f0000\""),
	// 	ID: to.Ptr("/subscriptions/8c3c936a-c09b-4de3-830b-3f5f244d72e9/resourceGroups/contosoithyperv/providers/Microsoft.Migrate/assessmentprojects/migrateprojectce73project/hypervcollectors/migrateprojectce73collector"),
	// 	Properties: &armmigrate.CollectorProperties{
	// 		AgentProperties: &armmigrate.CollectorAgentProperties{
	// 			ID: to.Ptr("d86c7d5a-2103-5157-bb20-9026b75e5de8"),
	// 			LastHeartbeatUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-08T09:43:59.057Z"); return t}()),
	// 			SpnDetails: &armmigrate.CollectorBodyAgentSpnProperties{
	// 				ApplicationID: to.Ptr("827f1053-44dc-439f-b832-05416dcce12b"),
	// 				Audience: to.Ptr("https://72f988bf-86f1-41af-91ab-2d7cd011db47/migrateprojectce73agentauthaadapp"),
	// 				Authority: to.Ptr("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 				ObjectID: to.Ptr("be75098e-c0fc-4ac4-98c7-282ebbcf8370"),
	// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			},
	// 			Version: to.Ptr("1.0.8.218"),
	// 		},
	// 		CreatedTimestamp: to.Ptr("2019-04-10T10:22:51.6271816Z"),
	// 		DiscoverySiteID: to.Ptr("/subscriptions/8c3c936a-c09b-4de3-830b-3f5f244d72e9/resourceGroups/ContosoITHyperV/providers/Microsoft.OffAzure/HyperVSites/migrateprojectce73site"),
	// 		UpdatedTimestamp: to.Ptr("2019-09-08T09:43:59.0573145Z"),
	// 	},
	// }
}
