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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/PutBackupInstance.json
func ExampleBackupInstancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		armdataprotection.BackupInstanceResource{
			Properties: &armdataprotection.BackupInstance{
				DataSourceInfo: &armdataprotection.Datasource{
					DatasourceType:   to.Ptr("<datasource-type>"),
					ObjectType:       to.Ptr("<object-type>"),
					ResourceID:       to.Ptr("<resource-id>"),
					ResourceLocation: to.Ptr("<resource-location>"),
					ResourceName:     to.Ptr("<resource-name>"),
					ResourceType:     to.Ptr("<resource-type>"),
					ResourceURI:      to.Ptr("<resource-uri>"),
				},
				DataSourceSetInfo: &armdataprotection.DatasourceSet{
					DatasourceType:   to.Ptr("<datasource-type>"),
					ObjectType:       to.Ptr("<object-type>"),
					ResourceID:       to.Ptr("<resource-id>"),
					ResourceLocation: to.Ptr("<resource-location>"),
					ResourceName:     to.Ptr("<resource-name>"),
					ResourceType:     to.Ptr("<resource-type>"),
					ResourceURI:      to.Ptr("<resource-uri>"),
				},
				DatasourceAuthCredentials: &armdataprotection.SecretStoreBasedAuthCredentials{
					ObjectType: to.Ptr("<object-type>"),
					SecretStoreResource: &armdataprotection.SecretStoreResource{
						SecretStoreType: to.Ptr(armdataprotection.SecretStoreTypeAzureKeyVault),
						URI:             to.Ptr("<uri>"),
					},
				},
				FriendlyName: to.Ptr("<friendly-name>"),
				ObjectType:   to.Ptr("<object-type>"),
				PolicyInfo: &armdataprotection.PolicyInfo{
					PolicyID: to.Ptr("<policy-id>"),
					PolicyParameters: &armdataprotection.PolicyParameters{
						DataStoreParametersList: []armdataprotection.DataStoreParametersClassification{
							&armdataprotection.AzureOperationalStoreParameters{
								DataStoreType:   to.Ptr(armdataprotection.DataStoreTypesOperationalStore),
								ObjectType:      to.Ptr("<object-type>"),
								ResourceGroupID: to.Ptr("<resource-group-id>"),
							}},
					},
				},
				ValidationType: to.Ptr(armdataprotection.ValidationTypeShallowValidation),
			},
		},
		&armdataprotection.BackupInstancesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataprotection%2Farmdataprotection%2Fv0.4.0/sdk/resourcemanager/dataprotection/armdataprotection/README.md) on how to add the SDK to your project and authenticate.
