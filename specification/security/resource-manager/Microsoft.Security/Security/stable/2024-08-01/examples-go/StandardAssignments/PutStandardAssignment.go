package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2024-08-01/StandardAssignments/PutStandardAssignment.json
func ExampleStandardAssignmentsClient_Create_putAnAuditStandardAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStandardAssignmentsClient().Create(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", "1f3afdf9-d0c9-4c3d-847f-89da613e70a8", armsecurity.StandardAssignment{
		Properties: &armsecurity.StandardAssignmentProperties{
			Description: to.Ptr("Set of policies monitored by Azure Security Center for cross cloud"),
			AssignedStandard: &armsecurity.AssignedStandardItem{
				ID: to.Ptr("/providers/Microsoft.Security/securityStandards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
			},
			DisplayName:    to.Ptr("ASC Default"),
			Effect:         to.Ptr(armsecurity.EffectAudit),
			ExcludedScopes: []*string{},
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
	// 		ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/standardAssignments/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
	// 		Properties: &armsecurity.StandardAssignmentProperties{
	// 			Description: to.Ptr("Set of policies monitored by Azure Security Center for cross cloud"),
	// 			AssignedStandard: &armsecurity.AssignedStandardItem{
	// 				ID: to.Ptr("/providers/Microsoft.Security/securityStandards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
	// 			},
	// 			DisplayName: to.Ptr("ASC Default"),
	// 			Effect: to.Ptr(armsecurity.EffectAudit),
	// 			ExcludedScopes: []*string{
	// 			},
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
