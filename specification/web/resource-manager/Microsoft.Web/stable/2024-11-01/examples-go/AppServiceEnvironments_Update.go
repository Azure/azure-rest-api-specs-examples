package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/AppServiceEnvironments_Update.json
func ExampleEnvironmentsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnvironmentsClient().Update(ctx, "test-rg", "test-ase", armappservice.EnvironmentPatchResource{
		Properties: &armappservice.Environment{
			FrontEndScaleFactor: to.Ptr[int32](20),
			VirtualNetwork: &armappservice.VirtualNetworkProfile{
				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-subnet/subnets/delegated"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EnvironmentResource = armappservice.EnvironmentResource{
	// 	Name: to.Ptr("test-ase"),
	// 	Type: to.Ptr("Microsoft.Web/hostingEnvironments"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase"),
	// 	Kind: to.Ptr("ASEV3"),
	// 	Location: to.Ptr("South Central US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armappservice.Environment{
	// 		DedicatedHostCount: to.Ptr[int32](0),
	// 		DNSSuffix: to.Ptr("test-ase.p.azurewebsites.net"),
	// 		FrontEndScaleFactor: to.Ptr[int32](15),
	// 		HasLinuxWorkers: to.Ptr(true),
	// 		InternalLoadBalancingMode: to.Ptr(armappservice.LoadBalancingModeNone),
	// 		IpsslAddressCount: to.Ptr[int32](0),
	// 		MaximumNumberOfMachines: to.Ptr[int32](250),
	// 		MultiSize: to.Ptr("Standard_D2d_v4"),
	// 		NetworkingConfiguration: &armappservice.AseV3NetworkingConfiguration{
	// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/configurations/networking"),
	// 			Properties: &armappservice.AseV3NetworkingConfigurationProperties{
	// 				AllowNewPrivateEndpointConnections: to.Ptr(false),
	// 				ExternalInboundIPAddresses: []*string{
	// 					to.Ptr("52.153.248.36")},
	// 					FtpEnabled: to.Ptr(false),
	// 					InternalInboundIPAddresses: []*string{
	// 					},
	// 					LinuxOutboundIPAddresses: []*string{
	// 						to.Ptr("20.88.241.56"),
	// 						to.Ptr("20.88.241.9")},
	// 						RemoteDebugEnabled: to.Ptr(false),
	// 						WindowsOutboundIPAddresses: []*string{
	// 							to.Ptr("20.88.241.56"),
	// 							to.Ptr("20.88.241.9")},
	// 						},
	// 					},
	// 					ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
	// 					Status: to.Ptr(armappservice.HostingEnvironmentStatusReady),
	// 					Suspended: to.Ptr(false),
	// 					VirtualNetwork: &armappservice.VirtualNetworkProfile{
	// 						Name: to.Ptr("delegated"),
	// 						Type: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
	// 						ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-subnet/subnets/delegated"),
	// 						Subnet: to.Ptr(""),
	// 					},
	// 					ZoneRedundant: to.Ptr(false),
	// 				},
	// 			}
}
