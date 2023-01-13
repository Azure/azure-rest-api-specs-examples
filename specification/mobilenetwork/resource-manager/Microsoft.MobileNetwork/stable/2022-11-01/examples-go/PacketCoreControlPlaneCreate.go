package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/PacketCoreControlPlaneCreate.json
func ExamplePacketCoreControlPlanesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewPacketCoreControlPlanesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "TestPacketCoreCP", armmobilenetwork.PacketCoreControlPlane{
		Location: to.Ptr("eastus"),
		Properties: &armmobilenetwork.PacketCoreControlPlanePropertiesFormat{
			ControlPlaneAccessInterface: &armmobilenetwork.InterfaceProperties{
				Name: to.Ptr("N2"),
			},
			CoreNetworkTechnology: to.Ptr(armmobilenetwork.CoreNetworkTypeFiveGC),
			LocalDiagnosticsAccess: &armmobilenetwork.LocalDiagnosticsAccessConfiguration{
				AuthenticationType: to.Ptr(armmobilenetwork.AuthenticationTypeAAD),
				HTTPSServerCertificate: &armmobilenetwork.HTTPSServerCertificate{
					CertificateURL: to.Ptr("https://contosovault.vault.azure.net/certificates/ingress"),
				},
			},
			Platform: &armmobilenetwork.PlatformConfiguration{
				Type: to.Ptr(armmobilenetwork.PlatformTypeAKSHCI),
				AzureStackEdgeDevice: &armmobilenetwork.AzureStackEdgeDeviceResourceID{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/TestAzureStackEdgeDevice"),
				},
				ConnectedCluster: &armmobilenetwork.ConnectedClusterResourceID{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/TestConnectedCluster"),
				},
				CustomLocation: &armmobilenetwork.CustomLocationResourceID{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/TestCustomLocation"),
				},
			},
			Sites: []*armmobilenetwork.SiteResourceID{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite"),
				}},
			SKU:     to.Ptr(armmobilenetwork.BillingSKUG0),
			UeMtu:   to.Ptr[int32](1600),
			Version: to.Ptr("0.2.0"),
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
