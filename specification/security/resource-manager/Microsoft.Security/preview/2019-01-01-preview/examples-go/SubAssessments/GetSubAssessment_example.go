package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/SubAssessments/GetSubAssessment_example.json
func ExampleSubAssessmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubAssessmentsClient().Get(ctx, "subscriptions/212f9889-769e-45ae-ab43-6da33674bd26/resourceGroups/DEMORG/providers/Microsoft.Compute/virtualMachines/vm2", "1195afff-c881-495e-9bc5-1486211ae03f", "95f7da9c-a2a4-1322-0758-fcd24ef09b85", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SubAssessment = armsecurity.SubAssessment{
	// 	Name: to.Ptr("95f7da9c-a2a4-1322-0758-fcd24ef09b85"),
	// 	Type: to.Ptr("Microsoft.Security/assessments/subAssessments"),
	// 	ID: to.Ptr("/subscriptions/212f9889-769e-45ae-ab43-6da33674bd26/resourceGroups/DEMORG/providers/Microsoft.Compute/virtualMachines/vm2/providers/Microsoft.Security/assessments/1195afff-c881-495e-9bc5-1486211ae03f/subassessments/95f7da9c-a2a4-1322-0758-fcd24ef09b85"),
	// 	Properties: &armsecurity.SubAssessmentProperties{
	// 		Description: to.Ptr("PuTTY ssh_agent_channel_data Function Integer Overflow Vulnerability"),
	// 		AdditionalData: &armsecurity.AdditionalData{
	// 			AssessedResourceType: to.Ptr(armsecurity.AssessedResourceTypeServerVulnerability),
	// 		},
	// 		Category: to.Ptr("Local"),
	// 		DisplayName: to.Ptr("PuTTY ssh_agent_channel_data Function Integer Overflow Vulnerability"),
	// 		ID: to.Ptr("370361"),
	// 		Impact: to.Ptr("Successful exploitation could allow remote attackers to have unspecified impact via a large length value in an agent protocol message."),
	// 		Remediation: to.Ptr("Customers are advised to upgrade toPuTTY 0.68 or later version in order to remediate this vulnerability."),
	// 		ResourceDetails: &armsecurity.AzureResourceDetails{
	// 			Source: to.Ptr(armsecurity.SourceAzure),
	// 			ID: to.Ptr("/subscriptions/212f9889-769e-45ae-ab43-6da33674bd26/resourceGroups/DEMORG/providers/Microsoft.Compute/virtualMachines/vm2"),
	// 		},
	// 		Status: &armsecurity.SubAssessmentStatus{
	// 			Code: to.Ptr(armsecurity.SubAssessmentStatusCodeUnhealthy),
	// 			Severity: to.Ptr(armsecurity.SeverityMedium),
	// 		},
	// 		TimeGenerated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-02T12:36:50.779Z"); return t}()),
	// 	},
	// }
}
