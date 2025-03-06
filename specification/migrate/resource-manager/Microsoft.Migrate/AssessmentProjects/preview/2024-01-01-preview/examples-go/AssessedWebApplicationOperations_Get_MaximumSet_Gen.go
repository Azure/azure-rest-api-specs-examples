package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AssessedWebApplicationOperations_Get_MaximumSet_Gen.json
func ExampleAssessedWebApplicationOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessedWebApplicationOperationsClient().Get(ctx, "rgaksswagger", "testproject", "testaksassessment", "testaksassessmentapp", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessedWebApplication = armmigrationassessment.AssessedWebApplication{
	// 	Name: to.Ptr("testaksassessmentapp"),
	// 	Type: to.Ptr("AssessedWebApplication"),
	// 	ID: to.Ptr("/subscriptions/D6F60DF4-CE70-4E39-8217-B8FBE7CA85AA/resourceGroups/rgaksswagger/providers/Microsoft.Migrate/assessmentprojects/testproject/aksAssessments/testaksassessment/assessedWebApps/testaksassessmentapp"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-07T06:51:24.108Z"); return t}()),
	// 		CreatedBy: to.Ptr("User"),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-07T06:51:24.108Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("User"),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	ETag: to.Ptr("00000000-0000-0000-a616-12d4724c01d9"),
	// 	Properties: &armmigrationassessment.AssessedWebApplicationProperties{
	// 		DiscoveryArmID: to.Ptr("/subscriptions/D6F60DF4-CE70-4E39-8217-B8FBE7CA85AA/resourcegroups/testrg/providers/microsoft.offazure/vmwaresites/testsite/webApps/webapp"),
	// 		DisplayName: to.Ptr("webapp"),
	// 		MachineArmID: to.Ptr("/subscriptions/D6F60DF4-CE70-4E39-8217-B8FBE7CA85AA/resourcegroups/testrg/providers/microsoft.offazure/vmwaresites/testsite/machines/machine"),
	// 		MachineDisplayName: to.Ptr("machine"),
	// 		WebAppType: to.Ptr(armmigrationassessment.WebAppTypeIIS),
	// 		WebServerArmID: to.Ptr("/subscriptions/D6F60DF4-CE70-4E39-8217-B8FBE7CA85AA/resourcegroups/testrg/providers/microsoft.offazure/vmwaresites/testsite/webServers/webserver"),
	// 		WebServerDisplayName: to.Ptr("webserver"),
	// 		RecommendationResult: &armmigrationassessment.RecommendationResult{
	// 			ApproxMonthlyCost: to.Ptr[float32](3),
	// 			ClusterName: to.Ptr("testaksassessment-cluster"),
	// 			Limit: &armmigrationassessment.ComputeResource{
	// 				Cores: to.Ptr[float32](18),
	// 				Memory: to.Ptr[float32](13),
	// 			},
	// 			NodePoolArmSKUName: to.Ptr("StandardDS_v2"),
	// 			NodePoolID: to.Ptr("testaksassessmentCostDetail"),
	// 			NodePoolName: to.Ptr("testaksassessmentCostDetail"),
	// 			OSType: to.Ptr(armmigrationassessment.OSTypeLinux),
	// 			Request: &armmigrationassessment.ComputeResource{
	// 				Cores: to.Ptr[float32](12),
	// 				Memory: to.Ptr[float32](28),
	// 			},
	// 		},
	// 		SuitabilityResult: &armmigrationassessment.SuitabilityResult{
	// 			FailedChecks: []*armmigrationassessment.Check{
	// 				{
	// 					Name: to.Ptr("check"),
	// 					Description: to.Ptr("check"),
	// 					Behavior: to.Ptr(armmigrationassessment.CheckResultIssue),
	// 					Cause: to.Ptr("check"),
	// 					MoreInfo: &armmigrationassessment.MoreInfo{
	// 						Title: to.Ptr("title"),
	// 						URL: to.Ptr("https://www.bing.com"),
	// 					},
	// 					Recommendation: to.Ptr("upgrade IIS version"),
	// 			}},
	// 			Readiness: to.Ptr(armmigrationassessment.ReadinessUnknown),
	// 			SuggestedMigrationTool: to.Ptr(armmigrationassessment.SuggestedMigrationToolNone),
	// 			Suitability: to.Ptr(armmigrationassessment.SuitabilityNotSuitable),
	// 		},
	// 	},
	// }
}
