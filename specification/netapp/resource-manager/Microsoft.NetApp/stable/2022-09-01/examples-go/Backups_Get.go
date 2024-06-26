package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-09-01/examples/Backups_Get.json
func ExampleBackupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupsClient().Get(ctx, "myRG", "account1", "pool1", "volume1", "backup1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Backup = armnetapp.Backup{
	// 	Name: to.Ptr("account1/pool1/volume1/backup1"),
	// 	Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/backups"),
	// 	ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/backups/backup1"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetapp.BackupProperties{
	// 		BackupType: to.Ptr(armnetapp.BackupTypeManual),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33Z"); return t}()),
	// 		Label: to.Ptr("myLabel"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Size: to.Ptr[int64](10011),
	// 	},
	// }
}
