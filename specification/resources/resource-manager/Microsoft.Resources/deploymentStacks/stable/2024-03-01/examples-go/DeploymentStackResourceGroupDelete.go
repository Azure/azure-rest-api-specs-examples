package armdeploymentstacks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentstacks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edacc3b43f9603efa119eabb6013d952d1dbe7d6/specification/resources/resource-manager/Microsoft.Resources/deploymentStacks/stable/2024-03-01/examples/DeploymentStackResourceGroupDelete.json
func ExampleClient_BeginDeleteAtResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentstacks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginDeleteAtResourceGroup(ctx, "deploymentStacksRG", "simpleDeploymentStack", &armdeploymentstacks.ClientBeginDeleteAtResourceGroupOptions{UnmanageActionResources: nil,
		UnmanageActionResourceGroups:   nil,
		UnmanageActionManagementGroups: nil,
		BypassStackOutOfSyncError:      nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
