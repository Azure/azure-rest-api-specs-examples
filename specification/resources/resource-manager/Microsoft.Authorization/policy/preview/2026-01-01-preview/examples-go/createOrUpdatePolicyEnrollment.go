package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/createOrUpdatePolicyEnrollment.json
func ExampleEnrollmentsClient_CreateOrUpdate_createOrUpdateAPolicyEnrollment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnrollmentsClient().CreateOrUpdate(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster", "DemoExpensiveVM", armpolicy.Enrollment{
		Properties: &armpolicy.EnrollmentProperties{
			PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
			PolicyDefinitionReferenceIDs: []*string{
				to.Ptr("Limit_Skus"),
			},
			DisplayName: to.Ptr("Enroll demo cluster"),
			Description: to.Ptr("Enroll demo cluster from limit sku"),
			Metadata: map[string]any{
				"reason": "Enrollment for a expensive VM demo",
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpolicy.EnrollmentsClientCreateOrUpdateResponse{
	// 	Enrollment: armpolicy.Enrollment{
	// 		Properties: &armpolicy.EnrollmentProperties{
	// 			PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
	// 			PolicyAssignmentInstanceID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			PolicyDefinitionReferenceIDs: []*string{
	// 				to.Ptr("Limit_Skus"),
	// 			},
	// 			DisplayName: to.Ptr("Enroll demo cluster"),
	// 			Description: to.Ptr("Enroll demo cluster from limit sku"),
	// 			Metadata: map[string]any{
	// 				"reason": "Enrollment for a expensive VM demo",
	// 			},
	// 		},
	// 		SystemData: &armpolicy.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.1075056Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.1075056Z"); return t}()),
	// 		},
	// 		ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster/providers/Microsoft.Authorization/policyEnrollments/DemoExpensiveVM"),
	// 		Type: to.Ptr("Microsoft.Authorization/policyEnrollments"),
	// 		ETag: to.Ptr(azcore.ETag("00000000-0000-0000-0000-000000000000")),
	// 		Name: to.Ptr("DemoExpensiveVM"),
	// 	},
	// }
}
