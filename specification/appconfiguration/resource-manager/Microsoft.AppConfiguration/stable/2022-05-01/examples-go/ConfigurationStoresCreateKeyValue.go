package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresCreateKeyValue.json
func ExampleKeyValuesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappconfiguration.NewKeyValuesClient("c80fb759-c965-4c6a-9110-9b2b2d038882", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"myResourceGroup",
		"contoso",
		"myKey$myLabel",
		&armappconfiguration.KeyValuesClientCreateOrUpdateOptions{KeyValueParameters: &armappconfiguration.KeyValue{
			Properties: &armappconfiguration.KeyValueProperties{
				Tags: map[string]*string{
					"tag1": to.Ptr("tagValue1"),
					"tag2": to.Ptr("tagValue2"),
				},
				Value: to.Ptr("myValue"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
