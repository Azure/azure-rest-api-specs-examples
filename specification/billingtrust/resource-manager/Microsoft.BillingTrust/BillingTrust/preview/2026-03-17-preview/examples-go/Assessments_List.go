package armbillingtrust_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingtrust/armbillingtrust"
)

// Generated from example definition: 2026-03-17-preview/Assessments_List.json
func ExampleAssessmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingtrust.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssessmentsClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/billing-edu-rg/providers/Microsoft.Program/educationEnrollments/default", nil)
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
		// page = armbillingtrust.AssessmentsClientListResponse{
		// 	AssessmentListResult: armbillingtrust.AssessmentListResult{
		// 		Value: []*armbillingtrust.Assessment{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/billing-edu-rg/providers/Microsoft.Program/educationEnrollments/default/providers/Microsoft.BillingTrust/assessments/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.BillingTrust/assessments"),
		// 				Properties: &armbillingtrust.AssessmentProperties{
		// 					AssessmentType: to.Ptr(armbillingtrust.AssessmentTypeEdu),
		// 					EvaluationState: to.Ptr(armbillingtrust.AssessmentStateSucceeded),
		// 					ProvisioningState: to.Ptr(armbillingtrust.ProvisioningStateSucceeded),
		// 				},
		// 				SystemData: &armbillingtrust.SystemData{
		// 					CreatedBy: to.Ptr("edu-admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armbillingtrust.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-17T10:00:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("edu-admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armbillingtrust.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-17T10:15:00.000Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
