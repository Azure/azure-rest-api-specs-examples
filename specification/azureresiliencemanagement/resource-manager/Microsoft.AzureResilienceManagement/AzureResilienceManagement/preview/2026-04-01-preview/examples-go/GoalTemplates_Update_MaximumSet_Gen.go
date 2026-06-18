package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/GoalTemplates_Update_MaximumSet_Gen.json
func ExampleGoalTemplatesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGoalTemplatesClient().BeginUpdate(ctx, "ipvrpvfcsfwltkmalhklsyg", "gt1", armresiliencemanagement.GoalTemplate{
		Properties: &armresiliencemanagement.GoalTemplateProperties{
			RequireHighAvailability:        to.Ptr(armresiliencemanagement.RequirementSelectedRequired),
			RequireDisasterRecovery:        to.Ptr(armresiliencemanagement.RequirementSelectedNotRequired),
			RegionalRecoveryPointObjective: to.Ptr("PT15M"),
			RegionalRecoveryTimeObjective:  to.Ptr("PT30M"),
			GoalType:                       to.Ptr(armresiliencemanagement.GoalTypeResiliency),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresiliencemanagement.GoalTemplatesClientUpdateResponse{
	// }
}
