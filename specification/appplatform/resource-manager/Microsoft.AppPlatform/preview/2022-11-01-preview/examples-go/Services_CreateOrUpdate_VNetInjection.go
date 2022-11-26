package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/Services_CreateOrUpdate_VNetInjection.json
func ExampleServicesClient_BeginCreateOrUpdate_servicesCreateOrUpdateVNetInjection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewServicesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", armappplatform.ServiceResource{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armappplatform.ClusterResourceProperties{
			NetworkProfile: &armappplatform.NetworkProfile{
				AppNetworkResourceGroup: to.Ptr("my-app-network-rg"),
				AppSubnetID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork/subnets/apps"),
				IngressConfig: &armappplatform.IngressConfig{
					ReadTimeoutInSeconds: to.Ptr[int32](300),
				},
				ServiceCidr:                        to.Ptr("10.8.0.0/16,10.244.0.0/16,10.245.0.1/16"),
				ServiceRuntimeNetworkResourceGroup: to.Ptr("my-service-runtime-network-rg"),
				ServiceRuntimeSubnetID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork/subnets/serviceRuntime"),
			},
			VnetAddons: &armappplatform.ServiceVNetAddons{
				LogStreamPublicEndpoint: to.Ptr(true),
			},
		},
		SKU: &armappplatform.SKU{
			Name: to.Ptr("S0"),
			Tier: to.Ptr("Standard"),
		},
	}, nil)
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
