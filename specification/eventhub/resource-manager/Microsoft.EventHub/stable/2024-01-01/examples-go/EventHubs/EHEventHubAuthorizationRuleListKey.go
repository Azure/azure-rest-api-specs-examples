package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubAuthorizationRuleListKey.json
func ExampleEventHubsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventHubsClient().ListKeys(ctx, "ArunMonocle", "sdk-namespace-960", "sdk-EventHub-532", "sdk-Authrules-2513", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessKeys = armeventhub.AccessKeys{
	// 	KeyName: to.Ptr("sdk-Authrules-2513"),
	// 	PrimaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-960.servicebus.windows-int.net/;SharedAccessKeyName=sdk-Authrules-2513;SharedAccessKey=############################################;EntityPath=sdk-EventHub-532"),
	// 	PrimaryKey: to.Ptr("############################################"),
	// 	SecondaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-960.servicebus.windows-int.net/;SharedAccessKeyName=sdk-Authrules-2513;SharedAccessKey=############################################;EntityPath=sdk-EventHub-532"),
	// 	SecondaryKey: to.Ptr("############################################"),
	// }
}
