package armprogramenrollment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/programenrollment/armprogramenrollment"
)

// Generated from example definition: 2026-03-01-preview/EduEnrollments_CreateOrUpdate.json
func ExampleEduEnrollmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprogramenrollment.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEduEnrollmentsClient().BeginCreateOrUpdate(ctx, "testrg", "default", armprogramenrollment.EduEnrollment{
		Location: to.Ptr("eastus"),
		Properties: &armprogramenrollment.EduEnrollmentProperties{
			Domains: []*armprogramenrollment.DomainGroup{
				{
					DomainNames: []*string{
						to.Ptr("university.edu"),
						to.Ptr("college.edu"),
					},
					TenantID: to.Ptr("00000000-0000-0000-0000-000000000001"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armprogramenrollment.EduEnrollmentsClientCreateOrUpdateResponse{
	// 	EduEnrollment: armprogramenrollment.EduEnrollment{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ProgramEnrollment/eduEnrollments/default"),
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.ProgramEnrollment/eduEnrollments"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armprogramenrollment.EduEnrollmentProperties{
	// 			ProvisioningState: to.Ptr(armprogramenrollment.ProvisioningStateSucceeded),
	// 			Domains: []*armprogramenrollment.DomainGroup{
	// 				{
	// 					DomainNames: []*string{
	// 						to.Ptr("university.edu"),
	// 						to.Ptr("college.edu"),
	// 					},
	// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 					State: to.Ptr(armprogramenrollment.DomainGroupStateSucceeded),
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armprogramenrollment.SystemData{
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armprogramenrollment.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armprogramenrollment.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
	// 		},
	// 	},
	// }
}
