package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/PacketCoreControlPlaneListBySubscription.json
func ExamplePacketCoreControlPlanesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPacketCoreControlPlanesClient().NewListBySubscriptionPager(nil)
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
		// page.PacketCoreControlPlaneListResult = armmobilenetwork.PacketCoreControlPlaneListResult{
		// 	Value: []*armmobilenetwork.PacketCoreControlPlane{
		// 		{
		// 			Name: to.Ptr("TestPacketCoreCP"),
		// 			Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlane"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP"),
		// 			SystemData: &armmobilenetwork.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armmobilenetwork.PacketCoreControlPlanePropertiesFormat{
		// 				ControlPlaneAccessInterface: &armmobilenetwork.InterfaceProperties{
		// 					Name: to.Ptr("N2"),
		// 				},
		// 				CoreNetworkTechnology: to.Ptr(armmobilenetwork.CoreNetworkTypeFiveGC),
		// 				Installation: &armmobilenetwork.Installation{
		// 					State: to.Ptr(armmobilenetwork.InstallationStateInstalled),
		// 				},
		// 				InstalledVersion: to.Ptr("0.2.0"),
		// 				LocalDiagnosticsAccess: &armmobilenetwork.LocalDiagnosticsAccessConfiguration{
		// 					AuthenticationType: to.Ptr(armmobilenetwork.AuthenticationTypePassword),
		// 					HTTPSServerCertificate: &armmobilenetwork.HTTPSServerCertificate{
		// 						CertificateURL: to.Ptr("https://contosovault.vault.azure.net/certificates/ingress"),
		// 						Provisioning: &armmobilenetwork.CertificateProvisioning{
		// 							State: to.Ptr(armmobilenetwork.CertificateProvisioningStateNotProvisioned),
		// 						},
		// 					},
		// 				},
		// 				Platform: &armmobilenetwork.PlatformConfiguration{
		// 					Type: to.Ptr(armmobilenetwork.PlatformTypeAKSHCI),
		// 					AzureStackEdgeDevice: &armmobilenetwork.AzureStackEdgeDeviceResourceID{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/TestAzureStackEdgeDevice"),
		// 					},
		// 					AzureStackEdgeDevices: []*armmobilenetwork.AzureStackEdgeDeviceResourceID{
		// 						{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/TestAzureStackEdgeDevice"),
		// 						},
		// 						{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/TestAzureStackEdgeDevice2"),
		// 					}},
		// 					ConnectedCluster: &armmobilenetwork.ConnectedClusterResourceID{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/TestConnectedCluster"),
		// 					},
		// 					CustomLocation: &armmobilenetwork.CustomLocationResourceID{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/TestCustomLocation"),
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
		// 				RollbackVersion: to.Ptr("0.1.0"),
		// 				Sites: []*armmobilenetwork.SiteResourceID{
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite"),
		// 				}},
		// 				SKU: to.Ptr(armmobilenetwork.BillingSKUG0),
		// 				UeMtu: to.Ptr[int32](1600),
		// 				Version: to.Ptr("0.2.0"),
		// 			},
		// 	}},
		// }
	}
}
