package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/AzureIaasVm/ProtectionIntent_Validate.json
func ExampleProtectionIntentClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewProtectionIntentClient("<subscription-id>", cred, nil)
	_, err = client.Validate(ctx,
		"<azure-region>",
		armrecoveryservicesbackup.PreValidateEnableBackupRequest{
			Properties:   to.StringPtr("<properties>"),
			ResourceID:   to.StringPtr("<resource-id>"),
			ResourceType: armrecoveryservicesbackup.DataSourceTypeVM.ToPtr(),
			VaultID:      to.StringPtr("<vault-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
