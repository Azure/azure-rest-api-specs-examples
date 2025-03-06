package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/GroupsOperations_ListByAssessmentProject_MaximumSet_Gen.json
func ExampleGroupsOperationsClient_NewListByAssessmentProjectPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGroupsOperationsClient().NewListByAssessmentProjectPager("ayagrawrg", "app18700project", nil)
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
		// page.GroupListResult = armmigrationassessment.GroupListResult{
		// 	Value: []*armmigrationassessment.Group{
		// 		{
		// 			Name: to.Ptr("kuchatur-test"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups"),
		// 			ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawrg/providers/Microsoft.Migrate/assessmentprojects/app18700project/groups/kuchatur-test"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:18.589Z"); return t}()),
		// 				CreatedBy: to.Ptr("bhjfiiwermbzqfoqxtxpjigj"),
		// 				CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:18.589Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("lrmhonmbodwva"),
		// 				LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmigrationassessment.GroupProperties{
		// 				ProvisioningState: to.Ptr(armmigrationassessment.ProvisioningStateSucceeded),
		// 				AreAssessmentsRunning: to.Ptr(true),
		// 				Assessments: []*string{
		// 					to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawrg/providers/Microsoft.Migrate/assessmentprojects/app18700project/groups/kuchatur-test/assessments/asm1")},
		// 					CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:29.161Z"); return t}()),
		// 					GroupStatus: to.Ptr(armmigrationassessment.GroupStatusCompleted),
		// 					GroupType: to.Ptr(armmigrationassessment.GroupTypeDefault),
		// 					MachineCount: to.Ptr[int32](20),
		// 					SupportedAssessmentTypes: []*armmigrationassessment.AssessmentType{
		// 						to.Ptr(armmigrationassessment.AssessmentTypeMachineAssessment),
		// 						to.Ptr(armmigrationassessment.AssessmentTypeAvsAssessment)},
		// 						UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:29.161Z"); return t}()),
		// 					},
		// 			}},
		// 		}
	}
}
