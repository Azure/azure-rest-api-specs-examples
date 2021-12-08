Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataprotection%2Farmdataprotection%2Fv0.1.0/sdk/resourcemanager/dataprotection/armdataprotection/README.md) on how to add the SDK to your project and authenticate.

```go
package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/BackupInstanceOperations/PutBackupInstance.json
func ExampleBackupInstancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<backup-instance-name>",
		armdataprotection.BackupInstanceResource{
			Properties: &armdataprotection.BackupInstance{
				DataSourceInfo: &armdataprotection.Datasource{
					DatasourceType:   to.StringPtr("<datasource-type>"),
					ObjectType:       to.StringPtr("<object-type>"),
					ResourceID:       to.StringPtr("<resource-id>"),
					ResourceLocation: to.StringPtr("<resource-location>"),
					ResourceName:     to.StringPtr("<resource-name>"),
					ResourceType:     to.StringPtr("<resource-type>"),
					ResourceURI:      to.StringPtr("<resource-uri>"),
				},
				DataSourceSetInfo: &armdataprotection.DatasourceSet{
					DatasourceType:   to.StringPtr("<datasource-type>"),
					ObjectType:       to.StringPtr("<object-type>"),
					ResourceID:       to.StringPtr("<resource-id>"),
					ResourceLocation: to.StringPtr("<resource-location>"),
					ResourceName:     to.StringPtr("<resource-name>"),
					ResourceType:     to.StringPtr("<resource-type>"),
					ResourceURI:      to.StringPtr("<resource-uri>"),
				},
				DatasourceAuthCredentials: &armdataprotection.SecretStoreBasedAuthCredentials{
					AuthCredentials: armdataprotection.AuthCredentials{
						ObjectType: to.StringPtr("<object-type>"),
					},
					SecretStoreResource: &armdataprotection.SecretStoreResource{
						SecretStoreType: armdataprotection.SecretStoreTypeAzureKeyVault.ToPtr(),
						URI:             to.StringPtr("<uri>"),
					},
				},
				FriendlyName: to.StringPtr("<friendly-name>"),
				ObjectType:   to.StringPtr("<object-type>"),
				PolicyInfo: &armdataprotection.PolicyInfo{
					PolicyID: to.StringPtr("<policy-id>"),
					PolicyParameters: &armdataprotection.PolicyParameters{
						DataStoreParametersList: []armdataprotection.DataStoreParametersClassification{
							&armdataprotection.AzureOperationalStoreParameters{
								DataStoreParameters: armdataprotection.DataStoreParameters{
									DataStoreType: armdataprotection.DataStoreTypesOperationalStore.ToPtr(),
									ObjectType:    to.StringPtr("<object-type>"),
								},
								ResourceGroupID: to.StringPtr("<resource-group-id>"),
							}},
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
	log.Printf("BackupInstanceResource.ID: %s\n", *res.ID)
}
```
