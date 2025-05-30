package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2024-12-01/ServiceConfigurationsPutWAC.json
func ExampleServiceConfigurationsClient_CreateOrupdate_serviceConfigurationsPutWac() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceConfigurationsClient().CreateOrupdate(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default", "default", "WAC", armhybridconnectivity.ServiceConfigurationResource{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.ServiceConfigurationsClientCreateOrupdateResponse{
	// 	ServiceConfigurationResource: &armhybridconnectivity.ServiceConfigurationResource{
	// 		Type: to.Ptr("Microsoft.HybridConnectivity/endpoints/serviceConfigurations"),
	// 		ID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default/serviceconfigurations/WAC"),
	// 		Properties: &armhybridconnectivity.ServiceConfigurationProperties{
	// 			Port: to.Ptr[int64](6516),
	// 			ProvisioningState: to.Ptr(armhybridconnectivity.ProvisioningStateSucceeded),
	// 			ServiceName: to.Ptr(armhybridconnectivity.ServiceNameWAC),
	// 		},
	// 	},
	// }
}
