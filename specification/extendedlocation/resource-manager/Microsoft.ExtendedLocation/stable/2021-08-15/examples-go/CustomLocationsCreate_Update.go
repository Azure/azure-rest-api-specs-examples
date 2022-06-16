package armextendedlocation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsCreate_Update.json
func ExampleCustomLocationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armextendedlocation.NewCustomLocationsClient("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testresourcegroup",
		"customLocation01",
		armextendedlocation.CustomLocation{
			Location: to.Ptr("West US"),
			Identity: &armextendedlocation.Identity{
				Type: to.Ptr(armextendedlocation.ResourceIdentityTypeSystemAssigned),
			},
			Properties: &armextendedlocation.CustomLocationProperties{
				Authentication: &armextendedlocation.CustomLocationPropertiesAuthentication{
					Type:  to.Ptr("KubeConfig"),
					Value: to.Ptr("<base64 KubeConfig>"),
				},
				ClusterExtensionIDs: []*string{
					to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedCluster/someCluster/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension")},
				DisplayName:    to.Ptr("customLocationLocation01"),
				HostResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01"),
				Namespace:      to.Ptr("namespace01"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
