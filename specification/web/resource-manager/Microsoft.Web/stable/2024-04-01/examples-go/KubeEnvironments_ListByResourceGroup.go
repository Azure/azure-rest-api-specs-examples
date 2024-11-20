package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/KubeEnvironments_ListByResourceGroup.json
func ExampleKubeEnvironmentsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKubeEnvironmentsClient().NewListByResourceGroupPager("examplerg", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.KubeEnvironmentCollection = armappservice.KubeEnvironmentCollection{
		// 	Value: []*armappservice.KubeEnvironment{
		// 		{
		// 			Name: to.Ptr("jlaw-demo1"),
		// 			Type: to.Ptr("Microsoft.Web/kubeEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.Web/kubeEnvironments/jlaw-demo1"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			ExtendedLocation: &armappservice.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			Properties: &armappservice.KubeEnvironmentProperties{
		// 				DefaultDomain: to.Ptr("jlaw-demo1.k4apps.io"),
		// 				InternalLoadBalancerEnabled: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armappservice.KubeEnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("20.42.33.145"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("demo1"),
		// 			Type: to.Ptr("Microsoft.Web/kubeEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.Web/kubeEnvironments/demo1"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			ExtendedLocation: &armappservice.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			Properties: &armappservice.KubeEnvironmentProperties{
		// 				DefaultDomain: to.Ptr("demo1.k4apps.io"),
		// 				InternalLoadBalancerEnabled: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armappservice.KubeEnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("52.142.21.61"),
		// 			},
		// 	}},
		// }
	}
}
