package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/hunts/CreateHunt.json
func ExampleHuntsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("bd794837-4d29-4647-9105-6339bfdb4e6a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHuntsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "163e7b2a-a2ec-4041-aaba-d878a38f265f", armsecurityinsights.Hunt{
		Properties: &armsecurityinsights.HuntProperties{
			Description: to.Ptr("Log4J Hunt Description"),
			AttackTactics: []*armsecurityinsights.AttackTactic{
				to.Ptr(armsecurityinsights.AttackTacticReconnaissance),
			},
			AttackTechniques: []*string{
				to.Ptr("T1595"),
			},
			DisplayName:      to.Ptr("Log4J new hunt"),
			HypothesisStatus: to.Ptr(armsecurityinsights.HypothesisStatusUnknown),
			Labels: []*string{
				to.Ptr("Label1"),
				to.Ptr("Label2"),
			},
			Owner: &armsecurityinsights.HuntOwner{
				ObjectID: to.Ptr("873b5263-5d34-4149-b356-ad341b01e123"),
			},
			Status: to.Ptr(armsecurityinsights.StatusNew),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.HuntsClientCreateOrUpdateResponse{
	// 	Hunt: armsecurityinsights.Hunt{
	// 		Name: to.Ptr("163e7b2a-a2ec-4041-aaba-d878a38f265f"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/hunts"),
	// 		Etag: to.Ptr("\"de00c408-0000-0c00-0000-62741e350000\""),
	// 		ID: to.Ptr("/subscriptions/bd794837-4d29-4647-9105-6339bfdb4e6a/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/hunts/163e7b2a-a2ec-4041-aaba-d878a38f265f"),
	// 		Properties: &armsecurityinsights.HuntProperties{
	// 			Description: to.Ptr("Log4J Hunt Description"),
	// 			AttackTactics: []*armsecurityinsights.AttackTactic{
	// 				to.Ptr(armsecurityinsights.AttackTacticReconnaissance),
	// 			},
	// 			AttackTechniques: []*string{
	// 				to.Ptr("T1595"),
	// 			},
	// 			DisplayName: to.Ptr("Log4J new hunt"),
	// 			HypothesisStatus: to.Ptr(armsecurityinsights.HypothesisStatusUnknown),
	// 			Labels: []*string{
	// 				to.Ptr("Label1"),
	// 				to.Ptr("Label2"),
	// 			},
	// 			Owner: &armsecurityinsights.HuntOwner{
	// 				Email: to.Ptr("testemail@microsoft.com"),
	// 				ObjectID: to.Ptr("873b5263-5d34-4149-b356-ad341b01e123"),
	// 				OwnerType: to.Ptr(armsecurityinsights.OwnerTypeUser),
	// 				UserPrincipalName: to.Ptr("John Doe"),
	// 			},
	// 			Status: to.Ptr(armsecurityinsights.StatusNew),
	// 		},
	// 	},
	// }
}
