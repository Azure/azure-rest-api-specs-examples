package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/DeletedManagedHsm_Get.json
func ExampleManagedHsmsClient_GetDeleted() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedHsmsClient().GetDeleted(ctx, "hsm1", "westus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeletedManagedHsm = armkeyvault.DeletedManagedHsm{
	// 	Name: to.Ptr("vault-agile-drawer-6404"),
	// 	Type: to.Ptr("Microsoft.KeyVault/deletedManagedHSMs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.KeyVault/locations/westus/deletedManagedHSMs/hsm1"),
	// 	Properties: &armkeyvault.DeletedManagedHsmProperties{
	// 		DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T00:00:59Z"); return t}()),
	// 		Location: to.Ptr("westus"),
	// 		MhsmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.KeyVault/managedHSMs/hsm1"),
	// 		PurgeProtectionEnabled: to.Ptr(true),
	// 		ScheduledPurgeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T00:00:59Z"); return t}()),
	// 		Tags: map[string]*string{
	// 			"Dept": to.Ptr("hsm"),
	// 			"Environment": to.Ptr("production"),
	// 		},
	// 	},
	// }
}
