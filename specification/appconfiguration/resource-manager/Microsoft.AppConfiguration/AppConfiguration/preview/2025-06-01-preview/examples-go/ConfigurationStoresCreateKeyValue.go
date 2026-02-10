package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v3"
)

// Generated from example definition: 2025-06-01-preview/ConfigurationStoresCreateKeyValue.json
func ExampleKeyValuesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("c80fb759-c965-4c6a-9110-9b2b2d038882", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKeyValuesClient().CreateOrUpdate(ctx, "myResourceGroup", "contoso", "myKey$myLabel", armappconfiguration.KeyValue{
		Properties: &armappconfiguration.KeyValueProperties{
			Tags: map[string]*string{
				"tag1": to.Ptr("tagValue1"),
				"tag2": to.Ptr("tagValue2"),
			},
			Value: to.Ptr("myValue"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappconfiguration.KeyValuesClientCreateOrUpdateResponse{
	// 	KeyValue: &armappconfiguration.KeyValue{
	// 		Name: to.Ptr("myKey$myLabel"),
	// 		Type: to.Ptr("Microsoft.AppConfiguration/configurationStores/keyValues"),
	// 		ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso/keyValues/myKey$myLabel"),
	// 		Properties: &armappconfiguration.KeyValueProperties{
	// 			ContentType: to.Ptr(""),
	// 			ETag: to.Ptr("IhDxoa8VkXxPsYsemBlxvV0d5fp"),
	// 			Key: to.Ptr("myKey"),
	// 			Label: to.Ptr("myLabel"),
	// 			LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-23T06:42:24+00:00"); return t}()),
	// 			Locked: to.Ptr(false),
	// 			Tags: map[string]*string{
	// 				"tag1": to.Ptr("tagValue1"),
	// 				"tag2": to.Ptr("tagValue2"),
	// 			},
	// 			Value: to.Ptr("myValue"),
	// 		},
	// 	},
	// }
}
