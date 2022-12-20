package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileServicesPut_EnableSecureSmbFeatures.json
func ExampleFileServicesClient_SetServiceProperties_putFileServicesEnableSecureSmbFeatures() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewFileServicesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.SetServiceProperties(ctx, "res4410", "sto8607", armstorage.FileServiceProperties{
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
	// TODO: use response item
	_ = res
}
