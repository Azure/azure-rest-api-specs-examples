package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/onboardingStates/CreateSentinelOnboardingState.json
func ExampleSentinelOnboardingStatesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSentinelOnboardingStatesClient().Create(ctx, "myRg", "myWorkspace", "default", &armsecurityinsights.SentinelOnboardingStatesClientCreateOptions{SentinelOnboardingStateParameter: &armsecurityinsights.SentinelOnboardingState{
		Properties: &armsecurityinsights.SentinelOnboardingStateProperties{
			CustomerManagedKey: to.Ptr(false),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SentinelOnboardingState = armsecurityinsights.SentinelOnboardingState{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/onboardingStates"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/onboardingStates/default"),
	// 	Properties: &armsecurityinsights.SentinelOnboardingStateProperties{
	// 		CustomerManagedKey: to.Ptr(false),
	// 	},
	// }
}
