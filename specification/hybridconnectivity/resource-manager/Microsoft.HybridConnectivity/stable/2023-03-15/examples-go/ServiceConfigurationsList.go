package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/ServiceConfigurationsList.json
func ExampleServiceConfigurationsClient_NewListByEndpointResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServiceConfigurationsClient().NewListByEndpointResourcePager("subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default", "default", nil)
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
		// page.ServiceConfigurationList = armhybridconnectivity.ServiceConfigurationList{
		// 	Value: []*armhybridconnectivity.ServiceConfigurationResource{
		// 		{
		// 			Type: to.Ptr("Microsoft.HybridConnectivity/endpoints/serviceConfigurations"),
		// 			ID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default/serviceconfigurations/SSH"),
		// 			Properties: &armhybridconnectivity.ServiceConfigurationProperties{
		// 				Port: to.Ptr[int64](22),
		// 				ServiceName: to.Ptr(armhybridconnectivity.ServiceNameSSH),
		// 			},
		// 		},
		// 		{
		// 			Type: to.Ptr("Microsoft.HybridConnectivity/endpoints/serviceConfigurations"),
		// 			ID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default/serviceconfigurations/WAC"),
		// 			Properties: &armhybridconnectivity.ServiceConfigurationProperties{
		// 				Port: to.Ptr[int64](6516),
		// 				ServiceName: to.Ptr(armhybridconnectivity.ServiceNameWAC),
		// 			},
		// 	}},
		// }
	}
}
