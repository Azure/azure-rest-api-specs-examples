package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Topics/SBTopicAuthorizationRuleListKey.json
func ExampleTopicsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopicsClient().ListKeys(ctx, "Default-ServiceBus-WestUS", "sdk-Namespace8408", "sdk-Topics2075", "sdk-Authrules5067", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessKeys = armservicebus.AccessKeys{
	// 	KeyName: to.Ptr("sdk-AuthRules-4310"),
	// 	PrimaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-6261.servicebus.windows-int.net/;SharedAccessKeyName=sdk-AuthRules-4310;SharedAccessKey=#############################################;EntityPath=sdk-Topics-1984"),
	// 	PrimaryKey: to.Ptr("#############################################"),
	// 	SecondaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-6261.servicebus.windows-int.net/;SharedAccessKeyName=sdk-AuthRules-4310;SharedAccessKey=#############################################;EntityPath=sdk-Topics-1984"),
	// 	SecondaryKey: to.Ptr("#############################################"),
	// }
}
