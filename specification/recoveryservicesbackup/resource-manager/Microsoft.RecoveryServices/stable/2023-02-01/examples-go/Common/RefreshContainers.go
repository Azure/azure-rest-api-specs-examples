package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/Common/RefreshContainers.json
func ExampleProtectionContainersClient_Refresh() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewProtectionContainersClient().Refresh(ctx, "NetSDKTestRsVault", "SwaggerTestRg", "Azure", &armrecoveryservicesbackup.ProtectionContainersClientRefreshOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
