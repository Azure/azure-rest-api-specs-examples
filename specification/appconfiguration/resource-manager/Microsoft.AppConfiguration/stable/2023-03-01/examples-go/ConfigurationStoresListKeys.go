package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/ConfigurationStoresListKeys.json
func ExampleConfigurationStoresClient_NewListKeysPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationStoresClient().NewListKeysPager("myResourceGroup", "contoso", &armappconfiguration.ConfigurationStoresClientListKeysOptions{SkipToken: nil})
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
		// page.APIKeyListResult = armappconfiguration.APIKeyListResult{
		// 	Value: []*armappconfiguration.APIKey{
		// 		{
		// 			Name: to.Ptr("Primary"),
		// 			ConnectionString: to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		// 			ID: to.Ptr("439AD01B4BE67DB1"),
		// 			LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:54.000Z"); return t}()),
		// 			ReadOnly: to.Ptr(false),
		// 			Value: to.Ptr("000000000000000000000000000000000000000000000000000000"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Secondary"),
		// 			ConnectionString: to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		// 			ID: to.Ptr("CB45E100456857B9"),
		// 			LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:54.000Z"); return t}()),
		// 			ReadOnly: to.Ptr(false),
		// 			Value: to.Ptr("000000000000000000000000000000000000000000000000000000"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Primary Read Only"),
		// 			ConnectionString: to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		// 			ID: to.Ptr("B3AC55B7E71431A9"),
		// 			LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:54.000Z"); return t}()),
		// 			ReadOnly: to.Ptr(true),
		// 			Value: to.Ptr("000000000000000000000000000000000000000000000000000000"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Secondary Read Only"),
		// 			ConnectionString: to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		// 			ID: to.Ptr("E2AF6A9A89DCC177"),
		// 			LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:54.000Z"); return t}()),
		// 			ReadOnly: to.Ptr(true),
		// 			Value: to.Ptr("000000000000000000000000000000000000000000000000000000"),
		// 	}},
		// }
	}
}
