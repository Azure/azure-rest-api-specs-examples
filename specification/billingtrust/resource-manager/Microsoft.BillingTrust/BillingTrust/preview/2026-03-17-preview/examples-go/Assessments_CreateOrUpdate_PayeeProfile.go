package armbillingtrust_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingtrust/armbillingtrust"
)

// Generated from example definition: 2026-03-17-preview/Assessments_CreateOrUpdate_PayeeProfile.json
func ExampleAssessmentsClient_BeginCreateOrUpdate_createOrUpdateThePayeeProfileAssessmentForABillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingtrust.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssessmentsClient().BeginCreateOrUpdate(ctx, "providers/Microsoft.Billing/billingAccounts/abc123:00000000-0000-0000-0000-000000000000_2019-05-31", armbillingtrust.Assessment{
		Properties: &armbillingtrust.AssessmentProperties{
			AssessmentType: to.Ptr(armbillingtrust.AssessmentTypePayeeProfile),
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
	// res = armbillingtrust.AssessmentsClientCreateOrUpdateResponse{
	// 	Assessment: armbillingtrust.Assessment{
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/abc123:00000000-0000-0000-0000-000000000000_2019-05-31/providers/Microsoft.BillingTrust/assessments/default"),
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.BillingTrust/assessments"),
	// 		Properties: &armbillingtrust.AssessmentProperties{
	// 			AssessmentType: to.Ptr(armbillingtrust.AssessmentTypePayeeProfile),
	// 			EvaluationState: to.Ptr(armbillingtrust.AssessmentStatePending),
	// 			ProvisioningState: to.Ptr(armbillingtrust.ProvisioningStateAccepted),
	// 		},
	// 		SystemData: &armbillingtrust.SystemData{
	// 			CreatedBy: to.Ptr("billing-admin@contoso.com"),
	// 			CreatedByType: to.Ptr(armbillingtrust.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-17T10:00:00.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("billing-admin@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armbillingtrust.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-17T10:00:00.000Z"); return t}()),
	// 		},
	// 	},
	// }
}
