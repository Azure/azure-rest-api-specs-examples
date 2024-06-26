package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/VMwareCollectors_Create.json
func ExampleVMwareCollectorsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVMwareCollectorsClient().Create(ctx, "abgoyal-westEurope", "abgoyalWEselfhostb72bproject", "PortalvCenterbc2fcollector", &armmigrate.VMwareCollectorsClientCreateOptions{CollectorBody: &armmigrate.VMwareCollector{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VMwareCollector = armmigrate.VMwareCollector{
	// 	Name: to.Ptr("PortalvCenterbc2fcollector"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/vmwarecollectors"),
	// 	ETag: to.Ptr("\"01003d32-0000-0d00-0000-5d74d2e50000\""),
	// 	ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westeurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject/vmwarecollectors/PortalvCenterbc2fcollector"),
	// 	Properties: &armmigrate.CollectorProperties{
	// 		AgentProperties: &armmigrate.CollectorAgentProperties{
	// 			ID: to.Ptr("75b0f71e-1272-4f29-a801-29cfa4b34a6e"),
	// 			LastHeartbeatUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-28T06:21:28.779Z"); return t}()),
	// 			SpnDetails: &armmigrate.CollectorBodyAgentSpnProperties{
	// 				ApplicationID: to.Ptr("fc717575-8173-4b21-92a5-658b655e613e"),
	// 				Audience: to.Ptr("https://72f988bf-86f1-41af-91ab-2d7cd011db47/PortalvCenterbc2fagentauthaadapp"),
	// 				Authority: to.Ptr("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 				ObjectID: to.Ptr("29d94f38-db94-4980-aec0-0cfd55ab1cd0"),
	// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			},
	// 			Version: to.Ptr("1.0.8.227"),
	// 		},
	// 		CreatedTimestamp: to.Ptr("2019-05-09T09:58:21.4988104Z"),
	// 		DiscoverySiteID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westEurope/providers/Microsoft.OffAzure/VMwareSites/PortalvCenterbc2fsite"),
	// 		UpdatedTimestamp: to.Ptr("2019-09-08T10:07:33.1996006Z"),
	// 	},
	// }
}
