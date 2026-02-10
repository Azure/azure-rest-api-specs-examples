package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v3"
)

// Generated from example definition: 2025-06-01-preview/DeletedConfigurationStoresList.json
func ExampleConfigurationStoresClient_NewListDeletedPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("c80fb759-c965-4c6a-9110-9b2b2d038882", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationStoresClient().NewListDeletedPager(nil)
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
		// page = armappconfiguration.ConfigurationStoresClientListDeletedResponse{
		// 	DeletedConfigurationStoreListResult: armappconfiguration.DeletedConfigurationStoreListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/providers/Microsoft.AppConfiguration/deletedConfigurationStores?api-version=2021-10-01-preview&%24skiptoken=HY3RaoMwAEX%2fRcbeYhJrnRXKYNWOuqpME0sfNcYui0Yxade19N8ne7hcDlzOvVuKX81eKKmt4G4dooLQwgqsL2NGHUDYV6o68Z4rY1e388RtNvRQn2vNJjEaMSgNvcbneMUcsKg8BFwft8DndQ0w9hu2QOiFLRs4TsNFNHzSMBFsGvTQGvuD%2f5bVuTOw4R03vPkH%2fVqNAlzm5SxfOwh7ACOA8POTlvPjILlaU1ke8jImOc23JCppQVfZnna0DXc4ISc3vSVuRo5zJE6%2bj25C3vwk2v2kEV2mMn7PyOc1DbtNGkonnzuLym1G400uI5QRZj0efw%3d%3d"),
		// 		Value: []*armappconfiguration.DeletedConfigurationStore{
		// 			{
		// 				Name: to.Ptr("contoso"),
		// 				Type: to.Ptr("Microsoft.AppConfiguration/deletedConfigurationStores"),
		// 				ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/providers/Microsoft.AppConfiguration/locations/westus/deletedConfigurationStores/contoso"),
		// 				Properties: &armappconfiguration.DeletedConfigurationStoreProperties{
		// 					ConfigurationStoreID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso"),
		// 					DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-01T00:00:59Z"); return t}()),
		// 					Location: to.Ptr("westus"),
		// 					PurgeProtectionEnabled: to.Ptr(true),
		// 					ScheduledPurgeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-01T00:00:59Z"); return t}()),
		// 					Tags: map[string]*string{
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
