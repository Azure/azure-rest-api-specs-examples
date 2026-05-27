package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2024-08-01/StandardAssignments/ListStandardAssignments.json
func ExampleStandardAssignmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStandardAssignmentsClient().NewListPager("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/gcpResourceGroup/providers/Microsoft.Security/securityConnectors/gcpconnector", nil)
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
		// page = armsecurity.StandardAssignmentsClientListResponse{
		// 	StandardAssignmentsList: armsecurity.StandardAssignmentsList{
		// 		Value: []*armsecurity.StandardAssignment{
		// 			{
		// 				Name: to.Ptr("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 				Type: to.Ptr("Microsoft.Security/standardAssignments"),
		// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/gcpResourceGroup/providers/Microsoft.Security/securityConnectors/gcpconnector/providers/Microsoft.Security/standardAssignments/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 				Properties: &armsecurity.StandardAssignmentProperties{
		// 					Description: to.Ptr("Set of policies monitored by Azure Security Center for cross cloud"),
		// 					AssignedStandard: &armsecurity.AssignedStandardItem{
		// 						ID: to.Ptr("/providers/Microsoft.Security/securityStandards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 					},
		// 					DisplayName: to.Ptr("ASC Default"),
		// 					Effect: to.Ptr(armsecurity.EffectAudit),
		// 					ExcludedScopes: []*string{
		// 					},
		// 					Metadata: &armsecurity.StandardAssignmentMetadata{
		// 						CreatedBy: to.Ptr("c23b5354-ff0a-4b2a-9f92-6f144effd936"),
		// 						CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T08:31:26.7993124Z"); return t}()),
		// 						LastUpdatedBy: to.Ptr("c23b5354-ff0a-4b2a-9f92-6f144effd936"),
		// 						LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T08:31:26.7993124Z"); return t}()),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 				Type: to.Ptr("Microsoft.Security/standardAssignments"),
		// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/gcpResourceGroup/providers/Microsoft.Security/securityConnectors/gcpconnector/providers/Microsoft.Security/standardAssignments/1f3afdf9-d0c9-4c3d-847f-89da613e70a9"),
		// 				Properties: &armsecurity.StandardAssignmentProperties{
		// 					Description: to.Ptr("Exemption description"),
		// 					AssignedStandard: &armsecurity.AssignedStandardItem{
		// 						ID: to.Ptr("/providers/Microsoft.Security/securityStandards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 					},
		// 					DisplayName: to.Ptr("Test exemption"),
		// 					Effect: to.Ptr(armsecurity.EffectExempt),
		// 					ExemptionData: &armsecurity.StandardAssignmentPropertiesExemptionData{
		// 						AssignedAssessment: &armsecurity.AssignedAssessmentItem{
		// 							AssessmentKey: to.Ptr("1195afff-c881-495e-9bc5-1486211ae03f"),
		// 						},
		// 						ExemptionCategory: to.Ptr(armsecurity.ExemptionCategoryWaiver),
		// 					},
		// 					ExpiresOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-01T19:50:47.083633Z"); return t}()),
		// 					Metadata: &armsecurity.StandardAssignmentMetadata{
		// 						CreatedBy: to.Ptr("c23b5354-ff0a-4b2a-9f92-6f144effd936"),
		// 						CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T08:31:26.7993124Z"); return t}()),
		// 						LastUpdatedBy: to.Ptr("c23b5354-ff0a-4b2a-9f92-6f144effd936"),
		// 						LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T08:31:26.7993124Z"); return t}()),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
