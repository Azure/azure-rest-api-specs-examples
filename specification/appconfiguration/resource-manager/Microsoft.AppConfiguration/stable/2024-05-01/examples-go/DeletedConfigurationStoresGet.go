package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d477c7caa09bf82e22c419be0a99d170552b5892/specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-05-01/examples/DeletedConfigurationStoresGet.json
func ExampleConfigurationStoresClient_GetDeleted() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationStoresClient().GetDeleted(ctx, "westus", "contoso", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeletedConfigurationStore = armappconfiguration.DeletedConfigurationStore{
	// 	Name: to.Ptr("contoso"),
	// 	Type: to.Ptr("Microsoft.AppConfiguration/deletedConfigurationStores"),
	// 	ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/providers/Microsoft.AppConfiguration/locations/westus/deletedConfigurationStores/contoso"),
	// 	Properties: &armappconfiguration.DeletedConfigurationStoreProperties{
	// 		ConfigurationStoreID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso"),
	// 		DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-01T00:00:59.000Z"); return t}()),
	// 		Location: to.Ptr("westus"),
	// 		PurgeProtectionEnabled: to.Ptr(true),
	// 		ScheduledPurgeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-01T00:00:59.000Z"); return t}()),
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
