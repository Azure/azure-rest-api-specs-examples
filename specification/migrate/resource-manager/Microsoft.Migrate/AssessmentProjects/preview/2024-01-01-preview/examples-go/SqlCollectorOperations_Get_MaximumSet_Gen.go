package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/SqlCollectorOperations_Get_MaximumSet_Gen.json
func ExampleSQLCollectorOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSQLCollectorOperationsClient().Get(ctx, "rgmigrate", "fci-test6904project", "fci-test0c1esqlsitecollector", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SQLCollector = armmigrationassessment.SQLCollector{
	// 	Name: to.Ptr("fci-test0c1esqlsitecollector"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/sqlcollectors"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.Migrate/assessmentprojects/fci-test6904project/sqlcollectors/fci-test0c1esqlsitecollector"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 	},
	// 	Properties: &armmigrationassessment.CollectorPropertiesBaseWithAgent{
	// 		AgentProperties: &armmigrationassessment.CollectorAgentPropertiesBase{
	// 			ID: to.Ptr("630da710-4d44-41f7-a189-72fe3db5502b-agent"),
	// 			SpnDetails: &armmigrationassessment.CollectorAgentSpnPropertiesBase{
	// 				ApplicationID: to.Ptr("db9c4c3d-477c-4d5a-817b-318276713565"),
	// 				Audience: to.Ptr("db9c4c3d-477c-4d5a-817b-318276713565"),
	// 				Authority: to.Ptr("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 				ObjectID: to.Ptr("e50236ad-ad07-47d4-af71-ed7b52d200d5"),
	// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			},
	// 		},
	// 		CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-22T09:50:37.106Z"); return t}()),
	// 		DiscoverySiteID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.OffAzure/MasterSites/fci-ankit-test6065mastersite/SqlSites/fci-ankit-test6065sqlsites"),
	// 		UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-16T12:17:52.918Z"); return t}()),
	// 	},
	// }
}
