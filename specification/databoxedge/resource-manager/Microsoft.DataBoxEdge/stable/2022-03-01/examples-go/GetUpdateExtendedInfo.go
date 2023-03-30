package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/GetUpdateExtendedInfo.json
func ExampleDevicesClient_UpdateExtendedInformation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().UpdateExtendedInformation(ctx, "testedgedevice", "GroupForEdgeAutomation", armdataboxedge.DeviceExtendedInfoPatch{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeviceExtendedInfo = armdataboxedge.DeviceExtendedInfo{
	// 	Properties: &armdataboxedge.DeviceExtendedInfoProperties{
	// 		ChannelIntegrityKeyName: to.Ptr("ase-cik-ab861822-21ea-4d31-96ec-89aa066f9a59"),
	// 		ChannelIntegrityKeyVersion: to.Ptr("95e9b619e65f433d82c9e1ead48214b9"),
	// 		ClientSecretStoreID: to.Ptr("/subscriptions/0d44739e-0563-474f-97e7-24a0cdb23b29/resourceGroups/arja-rg/providers/Microsoft.KeyVault/vaults/test-keyvault-ccy-523"),
	// 		ClientSecretStoreURL: to.Ptr("https://test-keyvault-ccy-523.vault.azure.net"),
	// 		DeviceSecrets: map[string]*armdataboxedge.Secret{
	// 			"BMCDefaultUserPassword": &armdataboxedge.Secret{
	// 			},
	// 			"HcsDataVolumeBitLockerExternalKey": &armdataboxedge.Secret{
	// 				KeyVaultID: to.Ptr("Id"),
	// 			},
	// 			"HcsInternalVolumeBitLockerExternalKey": &armdataboxedge.Secret{
	// 				KeyVaultID: to.Ptr("Id"),
	// 			},
	// 			"RotateKeyForDataVolumeBitlocker": &armdataboxedge.Secret{
	// 			},
	// 			"RotateKeysForSedDrivesSerialized": &armdataboxedge.Secret{
	// 			},
	// 			"SEDEncryptionExternalKey": &armdataboxedge.Secret{
	// 			},
	// 			"SEDEncryptionExternalKeyId": &armdataboxedge.Secret{
	// 			},
	// 			"SystemVolumeBitLockerRecoveryKey": &armdataboxedge.Secret{
	// 				KeyVaultID: to.Ptr("Id"),
	// 			},
	// 		},
	// 		KeyVaultSyncStatus: to.Ptr(armdataboxedge.KeyVaultSyncStatusKeyVaultSynced),
	// 	},
	// }
}
