package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2026-07-01-preview/TransparentDataEncryptions_ListByScheduler_MaximumSet_Gen.json
func ExampleTransparentDataEncryptionsClient_NewListBySchedulerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("EE9BD735-67CE-4A90-89C4-439D3F6A4C93", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTransparentDataEncryptionsClient().NewListBySchedulerPager("rgdurabletask", "testscheduler", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armdurabletask.TransparentDataEncryptionsClientListBySchedulerResponse{
		// 	TransparentDataEncryptionListResult: armdurabletask.TransparentDataEncryptionListResult{
		// 		Value: []*armdurabletask.TransparentDataEncryption{
		// 			{
		// 				Properties: &armdurabletask.TransparentDataEncryptionProperties{
		// 					ProvisioningState: to.Ptr(armdurabletask.ProvisioningStateSucceeded),
		// 					KeySource: to.Ptr(armdurabletask.TransparentDataEncryptionKeySourceCustomerManaged),
		// 					KeyVaultKeyURI: to.Ptr("https://microsoft.com/ap"),
		// 					AutoRotationEnabled: to.Ptr(true),
		// 				},
		// 				ID: to.Ptr("/subscriptions/EE9BD735-67CE-4A90-89C4-439D3F6A4C93/resourceGroups/rgdurabletask/providers/Microsoft.DurableTask/schedulers/testscheduler/transparentDataEncryptions/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.DurableTask/schedulers/transparentDataEncryptions"),
		// 				SystemData: &armdurabletask.SystemData{
		// 					CreatedBy: to.Ptr("tenmbevaunjzikxowqexrsx"),
		// 					CreatedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2024, time.April, 17, 15, 34, 17, 365000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("xfvdcegtj"),
		// 					LastModifiedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2024, time.April, 17, 15, 34, 17, 366000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
