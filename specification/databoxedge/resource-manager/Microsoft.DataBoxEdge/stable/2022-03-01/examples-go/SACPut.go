package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/SACPut.json
func ExampleStorageAccountCredentialsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageAccountCredentialsClient().BeginCreateOrUpdate(ctx, "testedgedevice", "sac1", "GroupForEdgeAutomation", armdataboxedge.StorageAccountCredential{
		Properties: &armdataboxedge.StorageAccountCredentialProperties{
			AccountKey: &armdataboxedge.AsymmetricEncryptedSecret{
				EncryptionAlgorithm:      to.Ptr(armdataboxedge.EncryptionAlgorithmAES256),
				EncryptionCertThumbprint: to.Ptr("2A9D8D6BE51574B5461230AEF02F162C5F01AD31"),
				Value:                    to.Ptr("lAeZEYi6rNP1/EyNaVUYmTSZEYyaIaWmwUsGwek0+xiZj54GM9Ue9/UA2ed/ClC03wuSit2XzM/cLRU5eYiFBwks23rGwiQOr3sruEL2a74EjPD050xYjA6M1I2hu/w2yjVHhn5j+DbXS4Xzi+rHHNZK3DgfDO3PkbECjPck+PbpSBjy9+6Mrjcld5DIZhUAeMlMHrFlg+WKRKB14o/og56u5/xX6WKlrMLEQ+y6E18dUwvWs2elTNoVO8PBE8SM/CfooX4AMNvaNdSObNBPdP+F6Lzc556nFNWXrBLRt0vC7s9qTiVRO4x/qCNaK/B4y7IqXMllwQFf4Np9UQ2ECA=="),
			},
			AccountType: to.Ptr(armdataboxedge.AccountTypeBlobStorage),
			Alias:       to.Ptr("sac1"),
			SSLStatus:   to.Ptr(armdataboxedge.SSLStatusDisabled),
			UserName:    to.Ptr("cisbvt"),
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
	// res.StorageAccountCredential = armdataboxedge.StorageAccountCredential{
	// 	Name: to.Ptr("sac1"),
	// 	Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/storageAccountCredentials"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccountCredentials/sac1"),
	// 	Properties: &armdataboxedge.StorageAccountCredentialProperties{
	// 		AccountType: to.Ptr(armdataboxedge.AccountTypeBlobStorage),
	// 		Alias: to.Ptr("sac1"),
	// 		SSLStatus: to.Ptr(armdataboxedge.SSLStatusDisabled),
	// 		UserName: to.Ptr("cisbvt"),
	// 	},
	// }
}
