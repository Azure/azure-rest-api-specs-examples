package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/CreateAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_CreateInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsMetadataClient().CreateInSubscription(ctx, "ca039e75-a276-4175-aebc-bcd41e4b14b7", armsecurity.AssessmentMetadataResponse{
		Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
			Description:    to.Ptr("Install an endpoint protection solution on your virtual machines scale sets, to protect them from threats and vulnerabilities."),
			AssessmentType: to.Ptr(armsecurity.AssessmentTypeCustomerManaged),
			Categories: []*armsecurity.Categories{
				to.Ptr(armsecurity.CategoriesCompute)},
			DisplayName:            to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
			ImplementationEffort:   to.Ptr(armsecurity.ImplementationEffortLow),
			RemediationDescription: to.Ptr("To install an endpoint protection solution: 1.  <a href=\"https://docs.microsoft.com/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-faq#how-do-i-turn-on-antimalware-in-my-virtual-machine-scale-set\">Follow the instructions in How do I turn on antimalware in my virtual machine scale set</a>"),
			Severity:               to.Ptr(armsecurity.SeverityMedium),
			Threats: []*armsecurity.Threats{
				to.Ptr(armsecurity.ThreatsDataExfiltration),
				to.Ptr(armsecurity.ThreatsDataSpillage),
				to.Ptr(armsecurity.ThreatsMaliciousInsider)},
			UserImpact: to.Ptr(armsecurity.UserImpactLow),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessmentMetadataResponse = armsecurity.AssessmentMetadataResponse{
	// 	Name: to.Ptr("ca039e75-a276-4175-aebc-bcd41e4b14b7"),
	// 	Type: to.Ptr("Microsoft.Security/assessmentMetadata"),
	// 	ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/ca039e75-a276-4175-aebc-bcd41e4b14b7"),
	// 	Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
	// 		Description: to.Ptr("Assessment that my organization created to view our security assessment in Azure Security Center"),
	// 		AssessmentType: to.Ptr(armsecurity.AssessmentTypeCustomerManaged),
	// 		Categories: []*armsecurity.Categories{
	// 			to.Ptr(armsecurity.CategoriesCompute)},
	// 			DisplayName: to.Ptr("My organization security assessment"),
	// 			ImplementationEffort: to.Ptr(armsecurity.ImplementationEffortLow),
	// 			RemediationDescription: to.Ptr("Fix it with these remediation instructions"),
	// 			Severity: to.Ptr(armsecurity.SeverityMedium),
	// 			Threats: []*armsecurity.Threats{
	// 				to.Ptr(armsecurity.ThreatsDataExfiltration),
	// 				to.Ptr(armsecurity.ThreatsDataSpillage),
	// 				to.Ptr(armsecurity.ThreatsMaliciousInsider)},
	// 				UserImpact: to.Ptr(armsecurity.UserImpactLow),
	// 			},
	// 		}
}
