package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubAuthorizationRuleListKey.json
func ExampleClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ListKeys(ctx, "5ktrial", "nh-sdk-ns", "nh-sdk-hub", "sdk-AuthRules-5800", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ResourceListKeys = armnotificationhubs.ResourceListKeys{
	// 	KeyName: to.Ptr("sdk-AuthRules-5800"),
	// 	PrimaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-7982.servicebus.windows-int.net/;SharedAccessKeyName=sdk-AuthRules-5800;SharedAccessKey=############################################;EntityPath=sdk-notificationHubs-2317"),
	// 	PrimaryKey: to.Ptr("############################################"),
	// 	SecondaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-7982.servicebus.windows-int.net/;SharedAccessKeyName=sdk-AuthRules-5800;SharedAccessKey=############################################;EntityPath=sdk-notificationHubs-2317"),
	// 	SecondaryKey: to.Ptr("############################################"),
	// }
}
