package armsecurity_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2024-08-01/StandardAssignments/PutExemptionStandardAssignment.json
func ExampleStandardAssignmentsClient_Create_putExemptionStandardAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStandardAssignmentsClient().Create(ctx, "subscriptions/212f9889-769e-45ae-ab43-6da33674bd26/resourceGroups/ANAT_TEST_RG/providers/Microsoft.Compute/virtualMachines/anatTestE2LA", "1f3afdf9-d0c9-4c3d-847f-89da613e70a8", armsecurity.StandardAssignment{
		Properties: &armsecurity.StandardAssignmentProperties{
			Description: to.Ptr("Exemption description"),
			AssignedStandard: &armsecurity.AssignedStandardItem{
				ID: to.Ptr("/providers/Microsoft.Security/securityStandards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
			},
			DisplayName: to.Ptr("Test exemption"),
			Effect:      to.Ptr(armsecurity.EffectExempt),
			ExemptionData: &armsecurity.StandardAssignmentPropertiesExemptionData{
				AssignedAssessment: &armsecurity.AssignedAssessmentItem{
					AssessmentKey: to.Ptr("1195afff-c881-495e-9bc5-1486211ae03f"),
				},
				ExemptionCategory: to.Ptr(armsecurity.ExemptionCategoryWaiver),
			},
			ExpiresOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-01T19:50:47.083633Z"); return t }()),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurity.StandardAssignmentsClientCreateResponse{
	// 	StandardAssignment: armsecurity.StandardAssignment{
	// 		Name: to.Ptr("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
	// 		Type: to.Ptr("Microsoft.Security/standardAssignments"),
	// 		ID: to.Ptr("/subscriptions/212f9889-769e-45ae-ab43-6da33674bd26/resourceGroups/ANAT_TEST_RG/providers/Microsoft.Compute/virtualMachines/anatTestE2LA/providers/Microsoft.Security/standardAssignments/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
	// 		Properties: &armsecurity.StandardAssignmentProperties{
	// 			Description: to.Ptr("Exemption description"),
	// 			AssignedStandard: &armsecurity.AssignedStandardItem{
	// 				ID: to.Ptr("/providers/Microsoft.Security/securityStandards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
	// 			},
	// 			DisplayName: to.Ptr("Test exemption"),
	// 			Effect: to.Ptr(armsecurity.EffectExempt),
	// 			ExemptionData: &armsecurity.StandardAssignmentPropertiesExemptionData{
	// 				AssignedAssessment: &armsecurity.AssignedAssessmentItem{
	// 					AssessmentKey: to.Ptr("1195afff-c881-495e-9bc5-1486211ae03f"),
	// 				},
	// 				ExemptionCategory: to.Ptr(armsecurity.ExemptionCategoryWaiver),
	// 			},
	// 			ExpiresOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-01T19:50:47.083633Z"); return t}()),
	// 			Metadata: &armsecurity.StandardAssignmentMetadata{
	// 				CreatedBy: to.Ptr("c23b5354-ff0a-4b2a-9f92-6f144effd936"),
	// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T08:31:26.7993124Z"); return t}()),
	// 				LastUpdatedBy: to.Ptr("c23b5354-ff0a-4b2a-9f92-6f144effd936"),
	// 				LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T08:31:26.7993124Z"); return t}()),
	// 			},
	// 		},
	// 	},
	// }
}
