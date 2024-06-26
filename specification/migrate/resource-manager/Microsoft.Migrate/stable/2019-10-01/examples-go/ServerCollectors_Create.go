package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/ServerCollectors_Create.json
func ExampleServerCollectorsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerCollectorsClient().Create(ctx, "pajindtest", "app11141project", "app23df4collector", &armmigrate.ServerCollectorsClientCreateOptions{CollectorBody: &armmigrate.ServerCollector{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerCollector = armmigrate.ServerCollector{
	// 	Name: to.Ptr("app23df4collector"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/servercollectors"),
	// 	ETag: to.Ptr("\"00000606-0000-0d00-0000-605999bf0000\""),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/pajindtest/providers/Microsoft.Migrate/assessmentprojects/app11141project/servercollectors/app23df4collector"),
	// 	Properties: &armmigrate.CollectorProperties{
	// 		AgentProperties: &armmigrate.CollectorAgentProperties{
	// 			ID: to.Ptr("dc984f5a-58a3-4f84-818c-a19febefa66a"),
	// 			LastHeartbeatUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-17T03:51:30.206Z"); return t}()),
	// 			SpnDetails: &armmigrate.CollectorBodyAgentSpnProperties{
	// 				ApplicationID: to.Ptr("ad9f701a-cc08-4421-b51f-b5762d58e9ba"),
	// 				Audience: to.Ptr("https://72f988bf-86f1-41af-91ab-2d7cd011db47/app23df4authandaccessaadapp"),
	// 				Authority: to.Ptr("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 				ObjectID: to.Ptr("b4975e42-9248-4a36-b99f-37eca377ea00"),
	// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			},
	// 			Version: to.Ptr("1.0.8.393"),
	// 		},
	// 		CreatedTimestamp: to.Ptr("2020-09-11T07:15:52.4361521Z"),
	// 		DiscoverySiteID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/pajindTest/providers/Microsoft.OffAzure/ServerSites/app21141site"),
	// 		UpdatedTimestamp: to.Ptr("2021-03-23T07:33:19.697297Z"),
	// 	},
	// }
}
