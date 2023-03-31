package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/340d577969b7bff5ad0488d79543314bc17daa50/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/PacketCoreControlPlaneCreate.json
func ExamplePacketCoreControlPlanesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPacketCoreControlPlanesClient().BeginCreateOrUpdate(ctx, "rg1", "TestPacketCoreCP", armmobilenetwork.PacketCoreControlPlane{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PacketCoreControlPlane = armmobilenetwork.PacketCoreControlPlane{
	// 	Name: to.Ptr("TestPacketCoreCP"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlane"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP"),
	// 	SystemData: &armmobilenetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armmobilenetwork.PacketCoreControlPlanePropertiesFormat{
	// 		ControlPlaneAccessInterface: &armmobilenetwork.InterfaceProperties{
	// 			Name: to.Ptr("N2"),
	// 		},
	// 		CoreNetworkTechnology: to.Ptr(armmobilenetwork.CoreNetworkTypeFiveGC),
	// 		Installation: &armmobilenetwork.Installation{
	// 			State: to.Ptr(armmobilenetwork.InstallationStateInstalled),
	// 		},
	// 		LocalDiagnosticsAccess: &armmobilenetwork.LocalDiagnosticsAccessConfiguration{
	// 			AuthenticationType: to.Ptr(armmobilenetwork.AuthenticationTypeAAD),
	// 			HTTPSServerCertificate: &armmobilenetwork.HTTPSServerCertificate{
	// 				CertificateURL: to.Ptr("https://contosovault.vault.azure.net/certificates/ingress"),
	// 				Provisioning: &armmobilenetwork.CertificateProvisioning{
	// 					State: to.Ptr(armmobilenetwork.CertificateProvisioningStateNotProvisioned),
	// 				},
	// 			},
	// 		},
	// 		Platform: &armmobilenetwork.PlatformConfiguration{
	// 			Type: to.Ptr(armmobilenetwork.PlatformTypeAKSHCI),
	// 			AzureStackEdgeDevice: &armmobilenetwork.AzureStackEdgeDeviceResourceID{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/TestAzureStackEdgeDevice"),
	// 			},
	// 			AzureStackEdgeDevices: []*armmobilenetwork.AzureStackEdgeDeviceResourceID{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/TestAzureStackEdgeDevice"),
	// 				},
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/TestAzureStackEdgeDevice2"),
	// 			}},
	// 			ConnectedCluster: &armmobilenetwork.ConnectedClusterResourceID{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/TestConnectedCluster"),
	// 			},
	// 			CustomLocation: &armmobilenetwork.CustomLocationResourceID{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/TestCustomLocation"),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
	// 		Sites: []*armmobilenetwork.SiteResourceID{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite"),
	// 		}},
	// 		SKU: to.Ptr(armmobilenetwork.BillingSKUG0),
	// 		UeMtu: to.Ptr[int32](1600),
	// 		Version: to.Ptr("0.2.0"),
	// 	},
	// }
}
