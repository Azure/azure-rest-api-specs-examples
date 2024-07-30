package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a651ba25cda4eec698a3a4e35f867ecc2681d126/specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/NodeTypePatchOperation_example.json
func ExampleNodeTypesClient_Update_patchANodeType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNodeTypesClient().Update(ctx, "resRg", "myCluster", "BE", armservicefabricmanagedclusters.NodeTypeUpdateParameters{
		Tags: map[string]*string{
			"a": to.Ptr("b"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NodeType = armservicefabricmanagedclusters.NodeType{
	// 	Name: to.Ptr("BE"),
	// 	Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/nodeTypes"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedClusters/myCluster/nodeTypes/BE"),
	// 	Tags: map[string]*string{
	// 		"a": to.Ptr("b"),
	// 	},
	// 	Properties: &armservicefabricmanagedclusters.NodeTypeProperties{
	// 		Capacities: map[string]*string{
	// 			"ClientConnections": to.Ptr("65536"),
	// 		},
	// 		DataDiskSizeGB: to.Ptr[int32](200),
	// 		DataDiskType: to.Ptr(armservicefabricmanagedclusters.DiskTypeStandardSSDLRS),
	// 		IsPrimary: to.Ptr(false),
	// 		IsStateless: to.Ptr(false),
	// 		PlacementProperties: map[string]*string{
	// 			"HasSSD": to.Ptr("true"),
	// 			"NodeColor": to.Ptr("green"),
	// 			"SomeProperty": to.Ptr("5"),
	// 		},
	// 		ProvisioningState: to.Ptr(armservicefabricmanagedclusters.ManagedResourceProvisioningStateSucceeded),
	// 		VMExtensions: []*armservicefabricmanagedclusters.VMSSExtension{
	// 			{
	// 				Name: to.Ptr("Microsoft.Azure.Geneva.GenevaMonitoring"),
	// 				Properties: &armservicefabricmanagedclusters.VMSSExtensionProperties{
	// 					Type: to.Ptr("GenevaMonitoring"),
	// 					AutoUpgradeMinorVersion: to.Ptr(true),
	// 					Publisher: to.Ptr("Microsoft.Azure.Geneva"),
	// 					Settings: map[string]any{
	// 					},
	// 					TypeHandlerVersion: to.Ptr("2.0"),
	// 				},
	// 		}},
	// 		VMImageOffer: to.Ptr("WindowsServer"),
	// 		VMImagePublisher: to.Ptr("MicrosoftWindowsServer"),
	// 		VMImageSKU: to.Ptr("2016-Datacenter-Server-Core"),
	// 		VMImageVersion: to.Ptr("latest"),
	// 		VMInstanceCount: to.Ptr[int32](10),
	// 		VMManagedIdentity: &armservicefabricmanagedclusters.VMManagedIdentity{
	// 			UserAssignedIdentities: []*string{
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity")},
	// 			},
	// 			VMSecrets: []*armservicefabricmanagedclusters.VaultSecretGroup{
	// 				{
	// 					SourceVault: &armservicefabricmanagedclusters.SubResource{
	// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.KeyVault/vaults/myVault"),
	// 					},
	// 					VaultCertificates: []*armservicefabricmanagedclusters.VaultCertificate{
	// 						{
	// 							CertificateStore: to.Ptr("My"),
	// 							CertificateURL: to.Ptr("https://myVault.vault.azure.net:443/secrets/myCert/ef1a31d39e1f46bca33def54b6cda54c"),
	// 					}},
	// 			}},
	// 			VMSize: to.Ptr("Standard_D3"),
	// 		},
	// 	}
}
