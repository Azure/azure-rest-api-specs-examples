package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2025-05-04/AssessmentsMetadata/GetAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_GetInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("0980887d-03d6-408c-9566-532f3456804e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsMetadataClient().GetInSubscription(ctx, "21300918-b2e3-0346-785f-c77ff57d243b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurity.AssessmentsMetadataClientGetInSubscriptionResponse{
	// 	AssessmentMetadataResponse: armsecurity.AssessmentMetadataResponse{
	// 		Name: to.Ptr("21300918-b2e3-0346-785f-c77ff57d243b"),
	// 		Type: to.Ptr("Microsoft.Security/assessmentMetadata"),
	// 		ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/21300918-b2e3-0346-785f-c77ff57d243b"),
	// 		Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
	// 			Description: to.Ptr("Install an endpoint protection solution on your virtual machines scale sets, to protect them from threats and vulnerabilities."),
	// 			AssessmentType: to.Ptr(armsecurity.AssessmentTypeBuiltIn),
	// 			Categories: []*armsecurity.Categories{
	// 				to.Ptr(armsecurity.CategoriesCompute),
	// 			},
	// 			DisplayName: to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
	// 			PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/26a828e1-e88f-464e-bbb3-c134a282b9de"),
	// 			RemediationDescription: to.Ptr("To install an endpoint protection solution: 1.  <a href=\"https://docs.microsoft.com/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-faq#how-do-i-turn-on-antimalware-in-my-virtual-machine-scale-set\">Follow the instructions in How do I turn on antimalware in my virtual machine scale set</a>"),
	// 			Severity: to.Ptr(armsecurity.SeverityMedium),
	// 			Threats: []*armsecurity.Threats{
	// 				to.Ptr(armsecurity.ThreatsDataExfiltration),
	// 				to.Ptr(armsecurity.ThreatsDataSpillage),
	// 				to.Ptr(armsecurity.ThreatsMaliciousInsider),
	// 			},
	// 		},
	// 	},
	// }
}
