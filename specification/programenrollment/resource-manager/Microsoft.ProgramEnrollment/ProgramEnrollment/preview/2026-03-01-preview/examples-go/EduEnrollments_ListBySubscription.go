package armprogramenrollment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/programenrollment/armprogramenrollment"
)

// Generated from example definition: 2026-03-01-preview/EduEnrollments_ListBySubscription.json
func ExampleEduEnrollmentsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprogramenrollment.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEduEnrollmentsClient().NewListBySubscriptionPager(nil)
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
		// page = armprogramenrollment.EduEnrollmentsClientListBySubscriptionResponse{
		// 	EduEnrollmentListResult: armprogramenrollment.EduEnrollmentListResult{
		// 		Value: []*armprogramenrollment.EduEnrollment{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ProgramEnrollment/eduEnrollments/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.ProgramEnrollment/eduEnrollments"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armprogramenrollment.EduEnrollmentProperties{
		// 					ProvisioningState: to.Ptr(armprogramenrollment.ProvisioningStateSucceeded),
		// 					Domains: []*armprogramenrollment.DomainGroup{
		// 						{
		// 							DomainNames: []*string{
		// 								to.Ptr("university.edu"),
		// 								to.Ptr("college.edu"),
		// 							},
		// 							TenantID: to.Ptr("00000000-0000-0000-0000-000000000001"),
		// 							State: to.Ptr(armprogramenrollment.DomainGroupStateSucceeded),
		// 						},
		// 					},
		// 				},
		// 				SystemData: &armprogramenrollment.SystemData{
		// 					CreatedBy: to.Ptr("user@example.com"),
		// 					CreatedByType: to.Ptr(armprogramenrollment.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@example.com"),
		// 					LastModifiedByType: to.Ptr(armprogramenrollment.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
