package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-10-01/Common/RefreshContainers_CrossSubscription.json
func ExampleProtectionContainersClient_Refresh_triggerCrossSubscriptionAzureWorkloadDiscovery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewProtectionContainersClient().Refresh(ctx, "NetSDKTestRsVault", "SwaggerTestRg", "Azure", &armrecoveryservicesbackup.ProtectionContainersClientRefreshOptions{
		Filter: to.Ptr("backupManagementType eq 'AzureWorkload' and containerSubscriptionId eq 'a76f4f58-8c04-4f53-9e68-4a698b0f43e4'")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
