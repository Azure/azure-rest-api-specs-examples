package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: 2025-05-01-preview/Topics/SBTopicAuthorizationRuleRegenerateKey.json
func ExampleTopicsClient_RegenerateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("e2f361f0-3b27-4503-a9cc-21cfba380093", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopicsClient().RegenerateKeys(ctx, "Default-ServiceBus-WestUS", "sdk-Namespace8408", "sdk-Topics2075", "sdk-Authrules5067", armservicebus.RegenerateAccessKeyParameters{
		KeyType: to.Ptr(armservicebus.KeyTypePrimaryKey),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicebus.TopicsClientRegenerateKeysResponse{
	// 	AccessKeys: armservicebus.AccessKeys{
	// 		KeyName: to.Ptr("sdk-AuthRules-4310"),
	// 		PrimaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-6261.servicebus.windows-int.net/;SharedAccessKeyName=sdk-AuthRules-4310;SharedAccessKey=#############################################;EntityPath=sdk-Topics-1984"),
	// 		PrimaryKey: to.Ptr("#############################################"),
	// 		SecondaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-6261.servicebus.windows-int.net/;SharedAccessKeyName=sdk-AuthRules-4310;SharedAccessKey=#############################################;EntityPath=sdk-Topics-1984"),
	// 		SecondaryKey: to.Ptr("#############################################"),
	// 	},
	// }
}
