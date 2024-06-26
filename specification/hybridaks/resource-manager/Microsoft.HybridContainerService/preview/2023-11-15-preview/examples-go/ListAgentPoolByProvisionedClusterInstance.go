package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2023-11-15-preview/examples/ListAgentPoolByProvisionedClusterInstance.json
func ExampleAgentPoolClient_ListByProvisionedCluster() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgentPoolClient().ListByProvisionedCluster(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgentPoolListResult = armhybridcontainerservice.AgentPoolListResult{
	// 	Value: []*armhybridcontainerservice.AgentPool{
	// 		{
	// 			Name: to.Ptr("test-hybridaksnodepool"),
	// 			Type: to.Ptr("Microsoft.HybridContainerService/provisionedClusterInstances/agentpools"),
	// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/agentPools/test-hybridaksnodepool"),
	// 			Location: to.Ptr("westus"),
	// 			Properties: &armhybridcontainerservice.AgentPoolProperties{
	// 				OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
	// 				ProvisioningState: to.Ptr(armhybridcontainerservice.ResourceProvisioningStateSucceeded),
	// 				Count: to.Ptr[int32](1),
	// 				VMSize: to.Ptr("Standard_A4_v2"),
	// 			},
	// 	}},
	// }
}
