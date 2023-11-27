package armextendedlocation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsPatch.json
func ExampleCustomLocationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomLocationsClient().Update(ctx, "testresourcegroup", "customLocation01", armextendedlocation.PatchableCustomLocations{
		Identity: &armextendedlocation.Identity{
			Type: to.Ptr(armextendedlocation.ResourceIdentityTypeSystemAssigned),
		},
		Properties: &armextendedlocation.CustomLocationProperties{
			ClusterExtensionIDs: []*string{
				to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension"),
				to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/barExtension")},
		},
		Tags: map[string]*string{
			"archv3": to.Ptr(""),
			"tier":   to.Ptr("testing"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomLocation = armextendedlocation.CustomLocation{
	// 	Name: to.Ptr("customLocation01"),
	// 	Type: to.Ptr("Microsoft.ExtendedLocation/customLocations"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/customLocation01"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"archv3": to.Ptr(""),
	// 		"tier": to.Ptr("testing"),
	// 	},
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
	// 			to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension"),
	// 			to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/barExtension")},
	// 			DisplayName: to.Ptr("customLocationLocation01"),
	// 			HostResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01"),
	// 			Namespace: to.Ptr("namespace01"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 		SystemData: &armextendedlocation.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-24T18:53:29.092Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-24T18:53:29.092Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
	// 		},
	// 	}
}
