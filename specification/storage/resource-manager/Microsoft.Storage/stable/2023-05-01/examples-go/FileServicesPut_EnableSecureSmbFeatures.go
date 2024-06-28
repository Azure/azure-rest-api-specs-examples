package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/FileServicesPut_EnableSecureSmbFeatures.json
func ExampleFileServicesClient_SetServiceProperties_putFileServicesEnableSecureSmbFeatures() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileServicesClient().SetServiceProperties(ctx, "res4410", "sto8607", armstorage.FileServiceProperties{
		FileServiceProperties: &armstorage.FileServicePropertiesProperties{
			ProtocolSettings: &armstorage.ProtocolSettings{
				Smb: &armstorage.SmbSetting{
					AuthenticationMethods:    to.Ptr("NTLMv2;Kerberos"),
					ChannelEncryption:        to.Ptr("AES-128-CCM;AES-128-GCM;AES-256-GCM"),
					KerberosTicketEncryption: to.Ptr("RC4-HMAC;AES-256"),
					Versions:                 to.Ptr("SMB2.1;SMB3.0;SMB3.1.1"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileServiceProperties = armstorage.FileServiceProperties{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res4410/providers/Microsoft.Storage/storageAccounts/sto8607/fileServices/default"),
	// 	FileServiceProperties: &armstorage.FileServicePropertiesProperties{
	// 		ProtocolSettings: &armstorage.ProtocolSettings{
	// 			Smb: &armstorage.SmbSetting{
	// 				AuthenticationMethods: to.Ptr("NTLMv2;Kerberos"),
	// 				ChannelEncryption: to.Ptr("AES-128-CCM;AES-128-GCM;AES-256-GCM"),
	// 				KerberosTicketEncryption: to.Ptr("RC4-HMAC;AES-256"),
	// 				Versions: to.Ptr("SMB2.1;SMB3.0;SMB3.1.1"),
	// 			},
	// 		},
	// 	},
	// 	SKU: &armstorage.SKU{
	// 		Name: to.Ptr(armstorage.SKUNamePremiumLRS),
	// 		Tier: to.Ptr(armstorage.SKUTierPremium),
	// 	},
	// }
}
