package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_Update.json
func ExamplePrivateCloudsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewPrivateCloudsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"group1",
		"cloud1",
		armavs.PrivateCloudUpdate{
			Identity: &armavs.PrivateCloudIdentity{
				Type: to.Ptr(armavs.ResourceIdentityTypeNone),
			},
			Properties: &armavs.PrivateCloudUpdateProperties{
				Encryption: &armavs.Encryption{
					KeyVaultProperties: &armavs.EncryptionKeyVaultProperties{
						KeyName:     to.Ptr("keyname1"),
						KeyVaultURL: to.Ptr("https://keyvault1-kmip-kvault.vault.azure.net/"),
						KeyVersion:  to.Ptr("ver1.0"),
					},
					Status: to.Ptr(armavs.EncryptionStateEnabled),
				},
				ManagementCluster: &armavs.ManagementCluster{
					ClusterSize: to.Ptr[int32](4),
				},
			},
		},
		nil)
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
