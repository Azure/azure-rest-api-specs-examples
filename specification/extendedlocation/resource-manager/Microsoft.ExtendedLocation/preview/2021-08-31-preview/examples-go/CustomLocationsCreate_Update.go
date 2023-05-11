package armextendedlocation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fb9c8e2ca33e9723c2b2fc849f627329067feb54/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/CustomLocationsCreate_Update.json
func ExampleCustomLocationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomLocationsClient().BeginCreateOrUpdate(ctx, "testresourcegroup", "customLocation01", armextendedlocation.CustomLocation{
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
	// res.CustomLocation = armextendedlocation.CustomLocation{
	// 	Name: to.Ptr("customLocation01"),
	// 	Type: to.Ptr("Microsoft.ExtendedLocation/customLocations"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/customLocation01"),
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armextendedlocation.Identity{
	// 		Type: to.Ptr(armextendedlocation.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		TenantID: to.Ptr("111111-1111-1111-1111-111111111111"),
	// 	},
	// 	Properties: &armextendedlocation.CustomLocationProperties{
	// 		Authentication: &armextendedlocation.CustomLocationPropertiesAuthentication{
	// 			Type: to.Ptr("KubeConfig"),
	// 		},
	// 		ClusterExtensionIDs: []*string{
	// 			to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension")},
	// 			DisplayName: to.Ptr("customLocationLocation01"),
	// 			HostResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01"),
	// 			Namespace: to.Ptr("namespace01"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 		SystemData: &armextendedlocation.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
	// 		},
	// 	}
}
