package armbillingtrust_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingtrust/armbillingtrust"
)

// Generated from example definition: 2026-03-17-preview/Rules_Update_BV.json
func ExampleRulesClient_Update_patchABusinessVerificationRuleWithExternalIdWhenActionRequiredResolvesAmbiguousMatchByDisambiguatingWithDuns() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingtrust.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRulesClient().Update(ctx, "providers/Microsoft.Billing/billingAccounts/abc123:00000000-0000-0000-0000-000000000000_2019-05-31", "Verify_Business", &armbillingtrust.BusinessVerificationRulePatchProperties{
		Kind: to.Ptr(armbillingtrust.RuleKindBusinessVerification),
		ExternalID: &armbillingtrust.ExternalID{
			Value: to.Ptr("123456789"),
			Type:  to.Ptr("DUNS"),
		},
		SupplementalDocuments: []*string{
			to.Ptr("https://trust.svc.cloud.microsoft/v1/documents/55555555-5555-5555-5555-555555555555"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbillingtrust.RulesClientUpdateResponse{
	// 	Rule: armbillingtrust.Rule{
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/abc123:00000000-0000-0000-0000-000000000000_2019-05-31/providers/Microsoft.BillingTrust/assessments/default/rules/Verify_Business"),
	// 		Name: to.Ptr("Verify_Business"),
	// 		Type: to.Ptr("Microsoft.BillingTrust/assessments/rules"),
	// 		Properties: &armbillingtrust.BusinessVerificationRuleProperties{
	// 			Kind: to.Ptr(armbillingtrust.RuleKindBusinessVerification),
	// 			EvaluationState: to.Ptr(armbillingtrust.RuleStateRunning),
	// 			ProvisioningState: to.Ptr(armbillingtrust.ProvisioningStateSucceeded),
	// 			SoldTo: &armbillingtrust.SoldTo{
	// 				CompanyName: to.Ptr("Contoso Ltd."),
	// 				FirstName: to.Ptr("Adele"),
	// 				LastName: to.Ptr("Vance"),
	// 				AddressLine1: to.Ptr("1 Microsoft Way"),
	// 				City: to.Ptr("Redmond"),
	// 				Region: to.Ptr("WA"),
	// 				PostalCode: to.Ptr("98052"),
	// 				Country: to.Ptr("US"),
	// 				Email: to.Ptr("billing@contoso.com"),
	// 				PhoneNumber: to.Ptr("+1-425-555-0100"),
	// 			},
	// 			RegistrationNumber: &armbillingtrust.RegistrationNumber{
	// 				Value: to.Ptr("12-3456789"),
	// 				Type: []*string{
	// 					to.Ptr("EIN"),
	// 				},
	// 			},
	// 			TaxIDs: []*armbillingtrust.TaxID{
	// 				{
	// 					Value: to.Ptr("12-3456789"),
	// 					Country: to.Ptr("US"),
	// 					Type: to.Ptr("EIN"),
	// 					Scope: to.Ptr("Federal"),
	// 					Status: to.Ptr(armbillingtrust.TaxIDStatusValid),
	// 				},
	// 			},
	// 			ExternalID: &armbillingtrust.ExternalID{
	// 				Value: to.Ptr("123456789"),
	// 				Type: to.Ptr("DUNS"),
	// 			},
	// 			SupplementalDocuments: []*string{
	// 				to.Ptr("https://trust.svc.cloud.microsoft/v1/documents/55555555-5555-5555-5555-555555555555"),
	// 			},
	// 		},
	// 		SystemData: &armbillingtrust.SystemData{
	// 			CreatedBy: to.Ptr("billing-admin@contoso.com"),
	// 			CreatedByType: to.Ptr(armbillingtrust.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-17T10:00:00.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("billing-admin@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armbillingtrust.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-17T11:30:00.000Z"); return t}()),
	// 		},
	// 	},
	// }
}
