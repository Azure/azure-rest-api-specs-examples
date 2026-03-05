package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v6"
)

// Generated from example definition: 2025-05-01/KubeEnvironments_Get.json
func ExampleKubeEnvironmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("8efdecc5-919e-44eb-b179-915dca89ebf9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKubeEnvironmentsClient().Get(ctx, "examplerg", "jlaw-demo1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappservice.KubeEnvironmentsClientGetResponse{
	// 	KubeEnvironment: &armappservice.KubeEnvironment{
	// 		Name: to.Ptr("jlaw-demo1"),
	// 		Type: to.Ptr("Microsoft.Web/kubeEnvironments"),
	// 		ExtendedLocation: &armappservice.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
	// 			Type: to.Ptr("customLocation"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.Web/kubeEnvironments/jlaw-demo1"),
	// 		Location: to.Ptr("North Central US"),
	// 		Properties: &armappservice.KubeEnvironmentProperties{
	// 			AksResourceID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ContainerService/managedClusters/jlaw-demo1"),
	// 			DefaultDomain: to.Ptr("jlaw-demo1.k4apps.io"),
	// 			InternalLoadBalancerEnabled: to.Ptr(false),
	// 			ProvisioningState: to.Ptr(armappservice.KubeEnvironmentProvisioningStateSucceeded),
	// 			StaticIP: to.Ptr("20.42.33.145"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
