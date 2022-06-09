```go
package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/SACPut.json
func ExampleStorageAccountCredentialsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewStorageAccountCredentialsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<device-name>",
		"<name>",
		"<resource-group-name>",
		armdataboxedge.StorageAccountCredential{
			Properties: &armdataboxedge.StorageAccountCredentialProperties{
				AccountKey: &armdataboxedge.AsymmetricEncryptedSecret{
					EncryptionAlgorithm:      armdataboxedge.EncryptionAlgorithm("AES256").ToPtr(),
					EncryptionCertThumbprint: to.StringPtr("<encryption-cert-thumbprint>"),
					Value:                    to.StringPtr("<value>"),
				},
				AccountType: armdataboxedge.AccountType("BlobStorage").ToPtr(),
				Alias:       to.StringPtr("<alias>"),
				SSLStatus:   armdataboxedge.SSLStatus("Disabled").ToPtr(),
				UserName:    to.StringPtr("<user-name>"),
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
	log.Printf("Response result: %#v\n", res.StorageAccountCredentialsClientCreateOrUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataboxedge%2Farmdataboxedge%2Fv0.2.0/sdk/resourcemanager/databoxedge/armdataboxedge/README.md) on how to add the SDK to your project and authenticate.
