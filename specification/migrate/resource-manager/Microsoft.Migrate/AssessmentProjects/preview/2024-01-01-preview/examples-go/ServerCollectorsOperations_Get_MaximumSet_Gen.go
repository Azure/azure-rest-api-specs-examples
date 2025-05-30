package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/ServerCollectorsOperations_Get_MaximumSet_Gen.json
func ExampleServerCollectorsOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerCollectorsOperationsClient().Get(ctx, "ayagrawRG", "app18700project", "walter389fcollector", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerCollector = armmigrationassessment.ServerCollector{
	// 	Name: to.Ptr("walter389fcollector"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/servercollectors"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawrg/providers/Microsoft.Migrate/assessmentprojects/app18700project/servercollectors/walter389fcollector"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 	},
	// 	Properties: &armmigrationassessment.CollectorPropertiesBaseWithAgent{
	// 		AgentProperties: &armmigrationassessment.CollectorAgentPropertiesBase{
	// 			ID: to.Ptr("498e4965-bbb1-47c2-8613-345baff9c509"),
	// 			SpnDetails: &armmigrationassessment.CollectorAgentSpnPropertiesBase{
	// 				ApplicationID: to.Ptr("65153d2f-9afb-44e8-b3ca-1369150b7354"),
	// 				Audience: to.Ptr("65153d2f-9afb-44e8-b3ca-1369150b7354"),
	// 				Authority: to.Ptr("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 				ObjectID: to.Ptr("ddde6f96-87c8-420b-9d4d-f16a5090519e"),
	// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			},
	// 		},
	// 		CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-29T13:35:13.362Z"); return t}()),
	// 		DiscoverySiteID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawRG/providers/Microsoft.OffAzure/ServerSites/walter7155site"),
	// 		UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-25T07:28:34.807Z"); return t}()),
	// 	},
	// }
}
