package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3751704f5318f1175875c94b66af769db917f2d3/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-01-01/examples/Common/ProtectedItem_Delete.json
func ExampleProtectedItemsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewProtectedItemsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "PySDKBackupTestRsVault", "PythonSDKBackupTestRg", "Azure", "iaasvmcontainer;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1", "vm;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
