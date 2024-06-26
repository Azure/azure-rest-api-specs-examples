package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/PutAgentPool.json
func ExampleAgentPoolClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgentPoolClient().BeginCreateOrUpdate(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster", "testnodepool", armhybridcontainerservice.AgentPool{
		Properties: &armhybridcontainerservice.AgentPoolProperties{
			NodeLabels: map[string]*string{
				"env":  to.Ptr("dev"),
				"goal": to.Ptr("test"),
			},
			NodeTaints: []*string{
				to.Ptr("env=prod:NoSchedule"),
				to.Ptr("sku=gpu:NoSchedule")},
			OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
			Count:  to.Ptr[int32](1),
			VMSize: to.Ptr("Standard_A4_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgentPool = armhybridcontainerservice.AgentPool{
	// 	Name: to.Ptr("testnodepool"),
	// 	Type: to.Ptr("Microsoft.HybridContainerService/provisionedClusterInstances/agentpools"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/agentPools/testnodepool"),
	// 	Properties: &armhybridcontainerservice.AgentPoolProperties{
	// 		NodeLabels: map[string]*string{
	// 			"env": to.Ptr("dev"),
	// 			"goal": to.Ptr("test"),
	// 		},
	// 		NodeTaints: []*string{
	// 			to.Ptr("env=prod:NoSchedule"),
	// 			to.Ptr("sku=gpu:NoSchedule")},
	// 			OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
	// 			ProvisioningState: to.Ptr(armhybridcontainerservice.ResourceProvisioningStateSucceeded),
	// 			Count: to.Ptr[int32](1),
	// 			VMSize: to.Ptr("Standard_A4_v2"),
	// 		},
	// 	}
}
