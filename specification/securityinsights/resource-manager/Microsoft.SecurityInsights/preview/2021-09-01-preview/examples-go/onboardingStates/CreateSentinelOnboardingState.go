package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/onboardingStates/CreateSentinelOnboardingState.json
func ExampleSentinelOnboardingStatesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewSentinelOnboardingStatesClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<sentinel-onboarding-state-name>",
		&armsecurityinsight.SentinelOnboardingStatesClientCreateOptions{SentinelOnboardingStateParameter: &armsecurityinsight.SentinelOnboardingState{
			Properties: &armsecurityinsight.SentinelOnboardingStateProperties{
				CustomerManagedKey: to.BoolPtr(false),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SentinelOnboardingStatesClientCreateResult)
}
