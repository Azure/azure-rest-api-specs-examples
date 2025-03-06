package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/WebAppCollectorOperations_ListByAssessmentProject_MaximumSet_Gen.json
func ExampleWebAppCollectorOperationsClient_NewListByAssessmentProjectPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebAppCollectorOperationsClient().NewListByAssessmentProjectPager("rgopenapi", "sumukk-ccy-bcs4557project", nil)
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
		// page.WebAppCollectorListResult = armmigrationassessment.WebAppCollectorListResult{
		// 	Value: []*armmigrationassessment.WebAppCollector{
		// 		{
		// 			Name: to.Ptr("sumukk-ccy-bcs4a93webappsitecollector"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/webappcollectors"),
		// 			ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sumukk-ccy-bcs/providers/Microsoft.Migrate/assessmentprojects/sumukk-ccy-bcs4557project/webappcollectors/sumukk-ccy-bcs4a93webappsitecollector"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:41:40.597Z"); return t}()),
		// 				CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:41:40.597Z"); return t}()),
		// 				LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmigrationassessment.WebAppCollectorPropertiesBaseWithAgent{
		// 				ProvisioningState: to.Ptr(armmigrationassessment.ProvisioningState2Succeeded),
		// 				AgentProperties: &armmigrationassessment.CollectorAgentPropertiesBase{
		// 					ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sumukk-ccy-bcs/providers/Microsoft.Migrate/assessmentprojects/sumukk-ccy-bcs4557project/webappcollectors/sumukk-ccy-bcs4a93webappsitecollector"),
		// 					LastHeartbeatUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:43:02.078Z"); return t}()),
		// 					SpnDetails: &armmigrationassessment.CollectorAgentSpnPropertiesBase{
		// 						ApplicationID: to.Ptr("2f70d5e8-7adc-4c64-910a-7031079efc6e"),
		// 						Audience: to.Ptr("2f70d5e8-7adc-4c64-910a-7031079efc6e"),
		// 						Authority: to.Ptr("https://login.microsoftonline.com/72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 						ObjectID: to.Ptr("2f70d5e8-7adc-4c64-910a-7031079efc6e"),
		// 						TenantID: to.Ptr("2f70d5e8-7adc-4c64-910a-7031079efc6e"),
		// 					},
		// 				},
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:43:02.078Z"); return t}()),
		// 				DiscoverySiteID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sumukk-ccy-bcs/providers/Microsoft.OffAzure/MasterSites/sumukk-ccy-bcs9880mastersite/WebAppSites/sumukk-ccy-bcs9880webappsites"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:43:02.078Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
