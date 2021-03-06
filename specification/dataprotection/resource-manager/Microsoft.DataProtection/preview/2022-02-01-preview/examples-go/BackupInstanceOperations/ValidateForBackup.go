package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/ValidateForBackup.json
func ExampleBackupInstancesClient_BeginValidateForBackup() {
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
	poller, err := client.BeginValidateForBackup(ctx,
		"<resource-group-name>",
		"<vault-name>",
		armdataprotection.ValidateForBackupRequest{
			BackupInstance: &armdataprotection.BackupInstance{
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
				},
			},
		},
		&armdataprotection.BackupInstancesClientBeginValidateForBackupOptions{ResumeToken: ""})
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
