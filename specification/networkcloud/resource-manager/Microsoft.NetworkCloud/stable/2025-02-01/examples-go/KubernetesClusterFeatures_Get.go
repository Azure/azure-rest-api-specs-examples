package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d0d3a9b4fe0fce880fded7a617e71f84406bacbd/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/KubernetesClusterFeatures_Get.json
func ExampleKubernetesClusterFeaturesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKubernetesClusterFeaturesClient().Get(ctx, "resourceGroupName", "kubernetesClusterName", "featureName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KubernetesClusterFeature = armnetworkcloud.KubernetesClusterFeature{
	// 	Name: to.Ptr("featureName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/kubernetesClusters/features"),
	// 	ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/kubernetesClusters/kubernetesClusterName/features/featureName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	Properties: &armnetworkcloud.KubernetesClusterFeatureProperties{
	// 		AvailabilityLifecycle: to.Ptr(armnetworkcloud.KubernetesClusterFeatureAvailabilityLifecycleGenerallyAvailable),
	// 		DetailedStatus: to.Ptr(armnetworkcloud.KubernetesClusterFeatureDetailedStatusInstalled),
	// 		DetailedStatusMessage: to.Ptr("The feature is successfully installed."),
	// 		Options: []*armnetworkcloud.StringKeyValuePair{
	// 			{
	// 				Key: to.Ptr("featureOptionName"),
	// 				Value: to.Ptr("featureOptionValue"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetworkcloud.KubernetesClusterFeatureProvisioningStateSucceeded),
	// 		Required: to.Ptr(armnetworkcloud.KubernetesClusterFeatureRequiredTrue),
	// 		Version: to.Ptr("1.2.3"),
	// 	},
	// }
}
