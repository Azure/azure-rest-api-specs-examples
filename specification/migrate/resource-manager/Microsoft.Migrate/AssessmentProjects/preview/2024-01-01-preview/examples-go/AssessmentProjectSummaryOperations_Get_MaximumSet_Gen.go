package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AssessmentProjectSummaryOperations_Get_MaximumSet_Gen.json
func ExampleAssessmentProjectSummaryOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentProjectSummaryOperationsClient().Get(ctx, "piyushapp1", "PiyushApp15328project", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessmentProjectSummary = armmigrationassessment.AssessmentProjectSummary{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/projectSummary"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/piyushapp1/providers/Microsoft.Migrate/assessmentprojects/PiyushApp15328project/projectSummary/default"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:18.589Z"); return t}()),
	// 		CreatedBy: to.Ptr("sakanwar"),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:18.589Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sakanwar"),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationassessment.AssessmentProjectSummaryProperties{
	// 		ErrorSummaryAffectedEntities: []*armmigrationassessment.ErrorSummary{
	// 			{
	// 				AssessmentType: to.Ptr(armmigrationassessment.AssessmentTypeMachineAssessment),
	// 				Count: to.Ptr[int32](1),
	// 			},
	// 			{
	// 				AssessmentType: to.Ptr(armmigrationassessment.AssessmentTypeAvsAssessment),
	// 				Count: to.Ptr[int32](2),
	// 		}},
	// 		LastAssessmentTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:23:46.901Z"); return t}()),
	// 		NumberOfAssessments: to.Ptr[int32](2),
	// 		NumberOfGroups: to.Ptr[int32](4),
	// 		NumberOfImportMachines: to.Ptr[int32](0),
	// 		NumberOfMachines: to.Ptr[int32](4),
	// 		NumberOfPrivateEndpointConnections: to.Ptr[int32](0),
	// 	},
	// }
}
