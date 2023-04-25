package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/GetAssessmentsMetadata_example.json
func ExampleAssessmentsMetadataClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsMetadataClient().Get(ctx, "21300918-b2e3-0346-785f-c77ff57d243b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessmentMetadataResponse = armsecurity.AssessmentMetadataResponse{
	// 	Name: to.Ptr("21300918-b2e3-0346-785f-c77ff57d243b"),
	// 	Type: to.Ptr("Microsoft.Security/assessmentMetadata"),
	// 	ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/21300918-b2e3-0346-785f-c77ff57d243b"),
	// 	Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
	// 		Description: to.Ptr("Install an endpoint protection solution on your virtual machines scale sets, to protect them from threats and vulnerabilities."),
	// 		AssessmentType: to.Ptr(armsecurity.AssessmentTypeBuiltIn),
	// 		Categories: []*armsecurity.Categories{
	// 			to.Ptr(armsecurity.CategoriesCompute)},
	// 			DisplayName: to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
	// 			ImplementationEffort: to.Ptr(armsecurity.ImplementationEffortLow),
	// 			PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/26a828e1-e88f-464e-bbb3-c134a282b9de"),
	// 			RemediationDescription: to.Ptr("To install an endpoint protection solution: 1.  <a href=\"https://docs.microsoft.com/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-faq#how-do-i-turn-on-antimalware-in-my-virtual-machine-scale-set\">Follow the instructions in How do I turn on antimalware in my virtual machine scale set</a>"),
	// 			Severity: to.Ptr(armsecurity.SeverityMedium),
	// 			Threats: []*armsecurity.Threats{
	// 				to.Ptr(armsecurity.ThreatsDataExfiltration),
	// 				to.Ptr(armsecurity.ThreatsDataSpillage),
	// 				to.Ptr(armsecurity.ThreatsMaliciousInsider)},
	// 				UserImpact: to.Ptr(armsecurity.UserImpactLow),
	// 				PlannedDeprecationDate: to.Ptr("03/2022"),
	// 				PublishDates: &armsecurity.AssessmentMetadataPropertiesResponsePublishDates{
	// 					GA: to.Ptr("06/01/2021"),
	// 					Public: to.Ptr("06/01/2021"),
	// 				},
	// 				Tactics: []*armsecurity.Tactics{
	// 					to.Ptr(armsecurity.TacticsCredentialAccess),
	// 					to.Ptr(armsecurity.TacticsPersistence),
	// 					to.Ptr(armsecurity.TacticsExecution),
	// 					to.Ptr(armsecurity.TacticsDefenseEvasion),
	// 					to.Ptr(armsecurity.TacticsCollection),
	// 					to.Ptr(armsecurity.TacticsDiscovery),
	// 					to.Ptr(armsecurity.TacticsPrivilegeEscalation)},
	// 					Techniques: []*armsecurity.Techniques{
	// 						to.Ptr(armsecurity.TechniquesObfuscatedFilesOrInformation),
	// 						to.Ptr(armsecurity.TechniquesIngressToolTransfer),
	// 						to.Ptr(armsecurity.TechniquesPhishing),
	// 						to.Ptr(armsecurity.TechniquesUserExecution)},
	// 					},
	// 				}
}
