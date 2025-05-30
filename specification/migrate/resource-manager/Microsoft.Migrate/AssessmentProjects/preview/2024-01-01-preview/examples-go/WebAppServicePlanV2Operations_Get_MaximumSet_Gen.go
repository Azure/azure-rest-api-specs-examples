package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/WebAppServicePlanV2Operations_Get_MaximumSet_Gen.json
func ExampleWebAppServicePlanV2OperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebAppServicePlanV2OperationsClient().Get(ctx, "rgopenapi", "sumukk-ccy-bcs4557project", "anraghun-selfhost-v2", "anraghun-selfhost-v2", "plan-001", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WebAppServicePlanV2 = armmigrationassessment.WebAppServicePlanV2{
	// 	Name: to.Ptr("plan-002-azureappservice"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/webAppAssessments/webAppServicePlans"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sumukk-ccy-bcs/providers/Microsoft.Migrate/assessmentprojects/sumukk-ccy-bcs4557project/groups/anraghun-selfhost-v2/webAppAssessments/anraghun-selfhost-v2/webAppServicePlans/plan-002-azureappservice"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:41:40.597Z"); return t}()),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:41:40.597Z"); return t}()),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationassessment.WebAppServicePlanV2Properties{
	// 		Cores: to.Ptr[int32](17),
	// 		CostComponents: []*armmigrationassessment.CostComponent{
	// 			{
	// 				Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlySecurityCost),
	// 				Value: to.Ptr[float32](16),
	// 		}},
	// 		CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:42:59.387Z"); return t}()),
	// 		MonthlyCost: to.Ptr[float32](20),
	// 		NumberOfWebApps: to.Ptr[int32](2),
	// 		RAM: to.Ptr[float64](18),
	// 		ScaleOutInstances: to.Ptr[int32](8),
	// 		Storage: to.Ptr[float64](25),
	// 		UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:42:59.387Z"); return t}()),
	// 		WebAppServicePlanName: to.Ptr("Plan-001"),
	// 		WebAppSKUName: to.Ptr("Isolated"),
	// 		WebAppSKUSize: to.Ptr("I1"),
	// 		WebAppTargetType: to.Ptr(armmigrationassessment.AzureWebAppTargetTypeAzureAppService),
	// 		WebAppType: to.Ptr(armmigrationassessment.WebAppTypeIIS),
	// 	},
	// }
}
