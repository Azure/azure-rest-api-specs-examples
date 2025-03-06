package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AssessedWebAppV2Operations_Get_MaximumSet_Gen.json
func ExampleAssessedWebAppV2OperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessedWebAppV2OperationsClient().Get(ctx, "rgopenapi", "sumukk-ccy-bcs4557project", "anraghun-selfhost-v2", "anraghun-selfhost-v2", "webapp1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessedWebAppV2 = armmigrationassessment.AssessedWebAppV2{
	// 	Name: to.Ptr("e1c05a87-93c7-4b01-bf19-a4db9bf791a8"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/webAppAssessments/assessedWebApps"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sumukk-ccy-bcs/providers/Microsoft.Migrate/assessmentprojects/sumukk-ccy-bcs4557project/groups/anraghun-selfhost-v2/webAppAssessments/anraghun-selfhost-v2/assessedWebApps/e1c05a87-93c7-4b01-bf19-a4db9bf791a8"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:41:40.597Z"); return t}()),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:41:40.597Z"); return t}()),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationassessment.AssessedWebAppV2Properties{
	// 		AppServicePlanName: to.Ptr("plan-001-azureappservice,plan-001-azureappservicecontainer"),
	// 		CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:42:51.551Z"); return t}()),
	// 		DiscoveredMachineID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourcegroups/sumukk-ccy-bcs/providers/microsoft.offazure/vmwaresites/sumukk-ccy-bcs4557site/machines/idclab-vcen67-fareast-corp-micr-49743448-0440-4c6b-8300-eec352b87e73_5037cd30-828f-f362-c62c-d1715b8ede4c"),
	// 		DiscoveredWebAppID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourcegroups/sumukk-ccy-bcs/providers/microsoft.offazure/mastersites/sumukk-ccy-bcs9880mastersite/webappsites/sumukk-ccy-bcs9880webappsites/iiswebapplications/aa096c1b26cd80dd76e547d61d722d04e1779a9c8d95ae1baa91bdb226eb5c3d-9e55bf838a799e058f7c0478ff1029d5224456ab0a1c30f3a8779ba7bb841e88"),
	// 		MachineName: to.Ptr("a404-r1w28r2-1"),
	// 		TargetSpecificResult: map[string]*armmigrationassessment.TargetSpecificResult{
	// 			"azureAppService": &armmigrationassessment.TargetSpecificResult{
	// 				AssessmentResult: &armmigrationassessment.AssessmentResult{
	// 					AppServicePlanName: to.Ptr("plan-001-azureappservice"),
	// 					SecuritySuitability: to.Ptr(armmigrationassessment.CloudSuitabilityUnknown),
	// 					Suitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
	// 					WebAppSKUName: to.Ptr("Isolated"),
	// 					WebAppSKUSize: to.Ptr("I1"),
	// 				},
	// 				MigrationIssues: []*armmigrationassessment.WebAppMigrationIssues{
	// 					{
	// 						IssueCategory: to.Ptr(armmigrationassessment.AzureWebAppSuitabilityIssueCategoryInfo),
	// 						IssueDescriptionList: []*string{
	// 							to.Ptr("SpecificUser (SharePoint - 80), SpecificUser (SecurityTokenServiceApplicationPool)")},
	// 							IssueID: to.Ptr("AppPoolIdentityCheck"),
	// 					}},
	// 				},
	// 			},
	// 			UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:42:51.551Z"); return t}()),
	// 			WebAppName: to.Ptr("100%WebSite"),
	// 			WebAppType: to.Ptr(armmigrationassessment.WebAppTypeIIS),
	// 			WebServerName: to.Ptr("IIS Server"),
	// 		},
	// 	}
}
