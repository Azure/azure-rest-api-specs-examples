package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2025-05-04/Assessments/GetAssessment_example.json
func ExampleAssessmentsClient_Get_getSecurityRecommendationTaskFromSecurityDataLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsClient().Get(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2", "21300918-b2e3-0346-785f-c77ff57d243b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurity.AssessmentsClientGetResponse{
	// 	AssessmentResponse: armsecurity.AssessmentResponse{
	// 		Name: to.Ptr("21300918-b2e3-0346-785f-c77ff57d243b"),
	// 		Type: to.Ptr("Microsoft.Security/assessments"),
	// 		ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2/providers/Microsoft.Security/assessments/21300918-b2e3-0346-785f-c77ff57d243b"),
	// 		Properties: &armsecurity.AssessmentPropertiesResponse{
	// 			AdditionalData: map[string]*string{
	// 				"linkedWorkspaceId": to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myLaWorkspace"),
	// 			},
	// 			DisplayName: to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
	// 			ResourceDetails: &armsecurity.AzureResourceDetails{
	// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2"),
	// 				Source: to.Ptr(armsecurity.SourceAzure),
	// 			},
	// 			Risk: &armsecurity.AssessmentPropertiesBaseRisk{
	// 				AttackPathsReferences: []*string{
	// 				},
	// 				IsContextualRisk: to.Ptr(true),
	// 				Level: to.Ptr(armsecurity.RiskLevelCritical),
	// 				Paths: []*armsecurity.AssessmentPropertiesBaseRiskPathsItem{
	// 					{
	// 						Edges: []*armsecurity.AssessmentPropertiesBaseRiskPathsItemEdgeItem{
	// 						},
	// 						ID: to.Ptr("14caae95be8a4e6da0b049ecb53bc880"),
	// 						Nodes: []*armsecurity.AssessmentPropertiesBaseRiskPathsItemNodesItem{
	// 							{
	// 								ID: to.Ptr("14caae95be8a4e6da0b049ecb53bc880"),
	// 								NodePropertiesLabel: []*string{
	// 									to.Ptr("exposedToInternet"),
	// 									to.Ptr("containsSensitiveData"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				RiskFactors: []*string{
	// 					to.Ptr("Exposure to the Internet"),
	// 					to.Ptr("Sensitive Data"),
	// 					to.Ptr("Critical Resource"),
	// 				},
	// 			},
	// 			Status: &armsecurity.AssessmentStatusResponse{
	// 				Description: to.Ptr("The effective policy for the assessment was evaluated to off - use Microsoft.Authorization/policyAssignments to turn this assessment on"),
	// 				Cause: to.Ptr("OffByPolicy"),
	// 				Code: to.Ptr(armsecurity.AssessmentStatusCodeNotApplicable),
	// 			},
	// 		},
	// 	},
	// }
}
