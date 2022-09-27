package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/CreateAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_CreateInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAssessmentsMetadataClient("0980887d-03d6-408c-9566-532f3456804e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateInSubscription(ctx, "ca039e75-a276-4175-aebc-bcd41e4b14b7", armsecurity.AssessmentMetadataResponse{
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
	// TODO: use response item
	_ = res
}
