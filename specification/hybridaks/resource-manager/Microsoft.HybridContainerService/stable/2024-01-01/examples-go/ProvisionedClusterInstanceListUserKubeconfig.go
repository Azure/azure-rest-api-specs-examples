package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/ProvisionedClusterInstanceListUserKubeconfig.json
func ExampleProvisionedClusterInstancesClient_BeginListUserKubeconfig() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProvisionedClusterInstancesClient().BeginListUserKubeconfig(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster", nil)
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
	// res.ListCredentialResponse = armhybridcontainerservice.ListCredentialResponse{
	// 	Name: to.Ptr("766ea16c-53c6-421e-9b7d-a8bea47285ed*36530D233A0F410A40772AE878D8E7A0B2223081048946AD3D40DE0268480FBE"),
	// 	ID: to.Ptr("/providers/Microsoft.HybridContainerService/locations/EASTUS/operationStatuses/766ea16c-53c6-421e-9b7d-a8bea47285ed*36530D233A0F410A40772AE878D8E7A0B2223081048946AD3D40DE0268480FBE"),
	// 	Properties: &armhybridcontainerservice.ListCredentialResponseProperties{
	// 		Kubeconfigs: []*armhybridcontainerservice.CredentialResult{
	// 			{
	// 				Name: to.Ptr("credentialName1"),
	// 				Value: []byte("Y3JlZGVudGlhbFZhbHVlMQ=="),
	// 		}},
	// 	},
	// 	ResourceID: to.Ptr("/subscriptions/921d26b3-c14d-4efc-b56e-93a2439e028c/resourceGroups/rg/providers/Microsoft.HybridContainerService/provisionedClusters/cluster-pc-1-24"),
	// 	Status: to.Ptr(armhybridcontainerservice.ResourceProvisioningStateSucceeded),
	// }
}
