package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomAssessmentAutomations/customAssessmentAutomationListBySubscription_example.json
func ExampleCustomAssessmentAutomationsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomAssessmentAutomationsClient().NewListBySubscriptionPager(nil)
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
		// page.CustomAssessmentAutomationsListResult = armsecurity.CustomAssessmentAutomationsListResult{
		// 	Value: []*armsecurity.CustomAssessmentAutomation{
		// 		{
		// 			Name: to.Ptr("MyCustomAssessmentAutomation1"),
		// 			Type: to.Ptr("Microsoft.Security/customAssessmentAutomations"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customAssessmentAutomations/MyCustomAssessmentAutomation1"),
		// 			Properties: &armsecurity.CustomAssessmentAutomationProperties{
		// 				Description: to.Ptr("organization passwords policy"),
		// 				AssessmentKey: to.Ptr("d5f442f7-7e77-4bcf-a450-a9c1b9a94eeb"),
		// 				CompressedQuery: to.Ptr("DQAKAEkAYQBtAF8ARwByAG8AdQBwAA0ACgB8ACAAZQB4AHQAZQBuAGQAIABIAGUAYQBsAHQAaABTAHQAYQB0AHUAcwAgAD0AIABpAGYAZgAoAHQAbwBzAHQAcgBpAG4AZwAoAFIAZQBjAG8AcgBkAC4AVQBzAGUAcgBOAGEAbQBlACkAIABjAG8AbgB0AGEAaQBuAHMAIAAnAHUAcwBlAHIAJwAsACAAJwBVAE4ASABFAEEATABUAEgAWQAnACwAIAAnAEgARQBBAEwAVABIAFkAJwApAA0ACgA="),
		// 				DisplayName: to.Ptr("Password Policy"),
		// 				RemediationDescription: to.Ptr("Change password policy to..."),
		// 				Severity: to.Ptr(armsecurity.SeverityEnumMedium),
		// 				SupportedCloud: to.Ptr(armsecurity.SupportedCloudEnumAWS),
		// 			},
		// 			SystemData: &armsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@contoso.com"),
		// 				CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("MyCustomAssessmentAutomation2"),
		// 			Type: to.Ptr("Microsoft.Security/customAssessmentAutomations"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customAssessmentAutomations/MyCustomAssessmentAutomation2"),
		// 			Properties: &armsecurity.CustomAssessmentAutomationProperties{
		// 				Description: to.Ptr("organization passwords policy"),
		// 				AssessmentKey: to.Ptr("fc1dbcd0-502c-4eab-9312-4014cfc8ea56"),
		// 				CompressedQuery: to.Ptr("Q29tcHV0ZV9OZXR3b3JrCnwgZXh0ZW5kIEhlYWx0aFN0YXR1cyA9ICdVTkhFQUxUSFkn"),
		// 				DisplayName: to.Ptr("Password Policy"),
		// 				RemediationDescription: to.Ptr("Change password policy to..."),
		// 				Severity: to.Ptr(armsecurity.SeverityEnumLow),
		// 				SupportedCloud: to.Ptr(armsecurity.SupportedCloudEnumGCP),
		// 			},
		// 			SystemData: &armsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:01:50.328Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@contoso.com"),
		// 				CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:01:50.328Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
