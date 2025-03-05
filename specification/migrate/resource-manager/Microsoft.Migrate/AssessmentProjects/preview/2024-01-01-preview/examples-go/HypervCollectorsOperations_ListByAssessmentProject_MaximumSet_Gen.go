package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/HypervCollectorsOperations_ListByAssessmentProject_MaximumSet_Gen.json
func ExampleHypervCollectorsOperationsClient_NewListByAssessmentProjectPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHypervCollectorsOperationsClient().NewListByAssessmentProjectPager("ayagrawRG", "app18700project", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.HypervCollectorListResult = armmigrationassessment.HypervCollectorListResult{
		// 	Value: []*armmigrationassessment.HypervCollector{
		// 		{
		// 			Name: to.Ptr("test-697cecollector"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/hypervcollectors"),
		// 			ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawrg/providers/Microsoft.Migrate/assessmentprojects/app18700project/hypervcollectors/test-697cecollector"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 			},
		// 			Properties: &armmigrationassessment.CollectorPropertiesBaseWithAgent{
		// 				AgentProperties: &armmigrationassessment.CollectorAgentPropertiesBase{
		// 					ID: to.Ptr("12f1d90f-b3fa-4926-8893-e56803a09af0"),
		// 					LastHeartbeatUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-07T14:25:35.708Z"); return t}()),
		// 					SpnDetails: &armmigrationassessment.CollectorAgentSpnPropertiesBase{
		// 						ApplicationID: to.Ptr("e3bd6eaa-980b-40ae-a30e-2a5069ba097c"),
		// 						Audience: to.Ptr("e3bd6eaa-980b-40ae-a30e-2a5069ba097c"),
		// 						Authority: to.Ptr("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 						ObjectID: to.Ptr("01b9f9e2-2d82-414c-adaa-09ce259b6b44"),
		// 						TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 					},
		// 					Version: to.Ptr("2.0.1993.19"),
		// 				},
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-06T17:35:52.750Z"); return t}()),
		// 				DiscoverySiteID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawRG/providers/Microsoft.OffAzure/HyperVSites/test-60527site"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-16T12:23:47.984Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
