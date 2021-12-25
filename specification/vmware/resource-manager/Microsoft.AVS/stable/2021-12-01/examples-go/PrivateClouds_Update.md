Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Favs%2Farmavs%2Fv0.1.0/sdk/resourcemanager/avs/armavs/README.md) on how to add the SDK to your project and authenticate.

```go
package armavs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_Update.json
func ExamplePrivateCloudsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewPrivateCloudsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		armavs.PrivateCloudUpdate{
			Identity: &armavs.PrivateCloudIdentity{
				Type: armavs.ResourceIdentityTypeNone.ToPtr(),
			},
			Properties: &armavs.PrivateCloudUpdateProperties{
				Encryption: &armavs.Encryption{
					KeyVaultProperties: &armavs.EncryptionKeyVaultProperties{
						KeyName:     to.StringPtr("<key-name>"),
						KeyVaultURL: to.StringPtr("<key-vault-url>"),
						KeyVersion:  to.StringPtr("<key-version>"),
					},
					Status: armavs.EncryptionStateEnabled.ToPtr(),
				},
				ManagementCluster: &armavs.ManagementCluster{
					CommonClusterProperties: armavs.CommonClusterProperties{
						ClusterSize: to.Int32Ptr(4),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PrivateCloud.ID: %s\n", *res.ID)
}
```
