package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/PrivateLinkResourceOperations_ListByAssessmentProject_MaximumSet_Gen.json
func ExamplePrivateLinkResourceOperationsClient_NewListByAssessmentProjectPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourceOperationsClient().NewListByAssessmentProjectPager("sakanwar", "sakanwar1204project", nil)
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
		// page.PrivateLinkResourceListResult = armmigrationassessment.PrivateLinkResourceListResult{
		// 	Value: []*armmigrationassessment.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("Default"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sakanwar/providers/Microsoft.Migrate/assessmentprojects/sakanwar1204project/privateLinkResources/Default"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:18.589Z"); return t}()),
		// 				CreatedBy: to.Ptr("sakanwar"),
		// 				CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:18.589Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sakanwar"),
		// 				LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmigrationassessment.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("Default"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("CollectorAgent")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.test.migration.windowsazure.com")},
		// 					},
		// 			}},
		// 		}
	}
}
