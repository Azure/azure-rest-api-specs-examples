package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a651ba25cda4eec698a3a4e35f867ecc2681d126/specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/NodeTypeListOperation_example.json
func ExampleNodeTypesClient_NewListByManagedClustersPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNodeTypesClient().NewListByManagedClustersPager("resRg", "myCluster", nil)
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
		// page.NodeTypeListResult = armservicefabricmanagedclusters.NodeTypeListResult{
		// 	Value: []*armservicefabricmanagedclusters.NodeType{
		// 		{
		// 			Name: to.Ptr("FE"),
		// 			Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/nodeTypes"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedClusters/myCluster/nodeTypes/FE"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicefabricmanagedclusters.NodeTypeProperties{
		// 				Capacities: map[string]*string{
		// 				},
		// 				DataDiskSizeGB: to.Ptr[int32](100),
		// 				DataDiskType: to.Ptr(armservicefabricmanagedclusters.DiskTypeStandardSSDLRS),
		// 				IsPrimary: to.Ptr(true),
		// 				IsStateless: to.Ptr(false),
		// 				PlacementProperties: map[string]*string{
		// 				},
		// 				VMImageOffer: to.Ptr("WindowsServer"),
		// 				VMImagePublisher: to.Ptr("MicrosoftWindowsServer"),
		// 				VMImageSKU: to.Ptr("2016-Datacenter"),
		// 				VMImageVersion: to.Ptr("latest"),
		// 				VMInstanceCount: to.Ptr[int32](5),
		// 				VMSize: to.Ptr("Standard_D2"),
		// 			},
		// 			SKU: &armservicefabricmanagedclusters.NodeTypeSKU{
		// 				Name: to.Ptr("Standard_P0"),
		// 				Capacity: to.Ptr[int32](5),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("BE"),
		// 			Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/nodeTypes"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedClusters/myCluster/nodeTypes/BE"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicefabricmanagedclusters.NodeTypeProperties{
		// 				Capacities: map[string]*string{
		// 				},
		// 				DataDiskSizeGB: to.Ptr[int32](200),
		// 				DataDiskType: to.Ptr(armservicefabricmanagedclusters.DiskTypePremiumLRS),
		// 				IsPrimary: to.Ptr(false),
		// 				IsStateless: to.Ptr(false),
		// 				PlacementProperties: map[string]*string{
		// 				},
		// 				VMExtensions: []*armservicefabricmanagedclusters.VMSSExtension{
		// 					{
		// 						Name: to.Ptr("Microsoft.Azure.Geneva.GenevaMonitoring"),
		// 						Properties: &armservicefabricmanagedclusters.VMSSExtensionProperties{
		// 							Type: to.Ptr("GenevaMonitoring"),
		// 							AutoUpgradeMinorVersion: to.Ptr(true),
		// 							Publisher: to.Ptr("Microsoft.Azure.Geneva"),
		// 							Settings: map[string]any{
		// 							},
		// 							TypeHandlerVersion: to.Ptr("2.0"),
		// 						},
		// 				}},
		// 				VMImageOffer: to.Ptr("WindowsServer"),
		// 				VMImagePublisher: to.Ptr("MicrosoftWindowsServer"),
		// 				VMImageSKU: to.Ptr("2016-Datacenter-Server-Core"),
		// 				VMImageVersion: to.Ptr("latest"),
		// 				VMInstanceCount: to.Ptr[int32](10),
		// 				VMSecrets: []*armservicefabricmanagedclusters.VaultSecretGroup{
		// 					{
		// 						SourceVault: &armservicefabricmanagedclusters.SubResource{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.KeyVault/vaults/myVault"),
		// 						},
		// 						VaultCertificates: []*armservicefabricmanagedclusters.VaultCertificate{
		// 							{
		// 								CertificateStore: to.Ptr("My"),
		// 								CertificateURL: to.Ptr("https://myVault.vault.azure.net:443/secrets/myCert/ef1a31d39e1f46bca33def54b6cda54c"),
		// 						}},
		// 				}},
		// 				VMSize: to.Ptr("Standard_DS3"),
		// 			},
		// 			SKU: &armservicefabricmanagedclusters.NodeTypeSKU{
		// 				Name: to.Ptr("Standard_S0"),
		// 				Capacity: to.Ptr[int32](10),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 	}},
		// }
	}
}
