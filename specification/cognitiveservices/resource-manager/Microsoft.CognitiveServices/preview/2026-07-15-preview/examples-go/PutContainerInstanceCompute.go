package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-07-15-preview/PutContainerInstanceCompute.json
func ExampleComputesClient_BeginCreateOrUpdate_putContainerInstanceCompute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewComputesClient().BeginCreateOrUpdate(ctx, "rgcognitiveservices", "myAccount", "myContainerInstance", armcognitiveservices.Compute{
		Properties: &armcognitiveservices.ContainerInstanceComputeProperties{
			ComputeType:            to.Ptr(armcognitiveservices.ComputeTypeContainerInstance),
			Location:               to.Ptr("eastus"),
			TargetClusterID:        to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/computes/myCluster"),
			ImageLink:              to.Ptr("mcr.microsoft.com/azureml/curated/pytorch-gpu:latest"),
			IdleTimeBeforeShutdown: to.Ptr("PT30M"),
			SSHSettings: &armcognitiveservices.SSHSettings{
				SSHPublicKey: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQ..."),
				AdminEnabled: to.Ptr(true),
			},
		},
		Identity: &armcognitiveservices.Identity{
			Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armcognitiveservices.UserAssignedIdentity{
				"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity": {},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
