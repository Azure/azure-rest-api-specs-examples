package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListRGDedicatedCloudServices.json
func ExampleDedicatedCloudServicesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedCloudServicesClient().NewListByResourceGroupPager("myResourceGroup", &armvmwarecloudsimple.DedicatedCloudServicesClientListByResourceGroupOptions{Filter: nil,
		Top:       nil,
		SkipToken: nil,
	})
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
		// page.DedicatedCloudServiceListResponse = armvmwarecloudsimple.DedicatedCloudServiceListResponse{
		// 	Value: []*armvmwarecloudsimple.DedicatedCloudService{
		// 		{
		// 			Name: to.Ptr("service-east"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/dedicatedCloudServices"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/service-east"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armvmwarecloudsimple.DedicatedCloudServiceProperties{
		// 				GatewaySubnet: to.Ptr("10.101.201.0/28"),
		// 				IsAccountOnboarded: to.Ptr(armvmwarecloudsimple.OnboardingStatusOnBoarded),
		// 				ServiceURL: to.Ptr("https://eastus-he.azure.cloudsimple.com"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("service-west"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/dedicatedCloudServices"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/service-west"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armvmwarecloudsimple.DedicatedCloudServiceProperties{
		// 				GatewaySubnet: to.Ptr("10.0.0.0/28"),
		// 				IsAccountOnboarded: to.Ptr(armvmwarecloudsimple.OnboardingStatusOnBoarded),
		// 				ServiceURL: to.Ptr("https://westus-he.azure.cloudsimple.com"),
		// 			},
		// 	}},
		// }
	}
}
