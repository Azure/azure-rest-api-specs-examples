package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresListKeyValues.json
func ExampleKeyValuesClient_NewListByConfigurationStorePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKeyValuesClient().NewListByConfigurationStorePager("myResourceGroup", "contoso", &armappconfiguration.KeyValuesClientListByConfigurationStoreOptions{SkipToken: nil})
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
		// page.KeyValueListResult = armappconfiguration.KeyValueListResult{
		// 	Value: []*armappconfiguration.KeyValue{
		// 		{
		// 			Name: to.Ptr("myKey$myLabel"),
		// 			Type: to.Ptr("Microsoft.AppConfiguration/configurationStores/keyValues"),
		// 			ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso/keyValues/myKey$myLabel"),
		// 			Properties: &armappconfiguration.KeyValueProperties{
		// 				ContentType: to.Ptr(""),
		// 				ETag: to.Ptr("IhDxoa8VkXxPsYsemBlxvV0d5fp"),
		// 				Key: to.Ptr("myKey"),
		// 				Label: to.Ptr("myLabel"),
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-23T06:42:24+00:00"); return t}()),
		// 				Locked: to.Ptr(false),
		// 				Tags: map[string]*string{
		// 					"tag1": to.Ptr("tagValue1"),
		// 					"tag2": to.Ptr("tagValue2"),
		// 				},
		// 				Value: to.Ptr("myValue"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myKey2$myLabel2"),
		// 			Type: to.Ptr("Microsoft.AppConfiguration/configurationStores/keyValues"),
		// 			ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso/keyValues/myKey2$myLabel2"),
		// 			Properties: &armappconfiguration.KeyValueProperties{
		// 				ContentType: to.Ptr(""),
		// 				ETag: to.Ptr("IfDxoa8VkXxPsYsemBlxvV0d5fp"),
		// 				Key: to.Ptr("myKey2"),
		// 				Label: to.Ptr("myLabel2"),
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-23T06:42:24+00:00"); return t}()),
		// 				Locked: to.Ptr(false),
		// 				Tags: map[string]*string{
		// 					"tag1": to.Ptr("tagValue1"),
		// 					"tag2": to.Ptr("tagValue2"),
		// 				},
		// 				Value: to.Ptr("myValue"),
		// 			},
		// 	}},
		// }
	}
}
