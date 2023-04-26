package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/GetAssessmentWithExpand_example.json
func ExampleAssessmentsClient_Get_getSecurityRecommendationTaskFromSecurityDataLocationWithExpandParameter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsClient().Get(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2", "21300918-b2e3-0346-785f-c77ff57d243b", &armsecurity.AssessmentsClientGetOptions{Expand: to.Ptr(armsecurity.ExpandEnumLinks)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessmentResponse = armsecurity.AssessmentResponse{
	// 	Name: to.Ptr("21300918-b2e3-0346-785f-c77ff57d243b"),
	// 	Type: to.Ptr("Microsoft.Security/assessments"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2/providers/Microsoft.Security/assessments/21300918-b2e3-0346-785f-c77ff57d243b"),
	// 	Properties: &armsecurity.AssessmentPropertiesResponse{
	// 		AdditionalData: map[string]*string{
	// 			"linkedWorkspaceId": to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myLaWorkspace"),
	// 		},
	// 		DisplayName: to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
	// 		Links: &armsecurity.AssessmentLinks{
	// 			AzurePortalURI: to.Ptr("https://www.portal.azure.com/?fea#blade/Microsoft_Azure_Security/RecommendationsBlade/assessmentKey/21300918-b2e3-0346-785f-c77ff57d243b"),
	// 		},
	// 		ResourceDetails: &armsecurity.AzureResourceDetails{
	// 			Source: to.Ptr(armsecurity.SourceAzure),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2"),
	// 		},
	// 		Status: &armsecurity.AssessmentStatusResponse{
	// 			Description: to.Ptr("The effective policy for the assessment was evaluated to off - use Microsoft.Authorization/policyAssignments to turn this assessment on"),
	// 			Cause: to.Ptr("OffByPolicy"),
	// 			Code: to.Ptr(armsecurity.AssessmentStatusCodeNotApplicable),
	// 			FirstEvaluationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-12T09:07:18.6759138Z"); return t}()),
	// 			StatusChangeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-12T09:07:18.6759138Z"); return t}()),
	// 		},
	// 	},
	// }
}
