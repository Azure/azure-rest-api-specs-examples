package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub/v2"
)

// Generated from example definition: 2025-05-01-preview/NameSpaces/EHNameSpaceAuthorizationRuleRegenerateKey.json
func ExampleNamespacesClient_RegenerateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().RegenerateKeys(ctx, "ArunMonocle", "sdk-Namespace-8980", "sdk-Authrules-8929", armeventhub.RegenerateAccessKeyParameters{
		KeyType: to.Ptr(armeventhub.KeyTypePrimaryKey),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armeventhub.NamespacesClientRegenerateKeysResponse{
	// 	AccessKeys: armeventhub.AccessKeys{
	// 		KeyName: to.Ptr("sdk-Authrules-8929"),
	// 		PrimaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-8980.servicebus.windows-int.net/;SharedAccessKeyName=sdk-Authrules-8929;SharedAccessKey=############################################"),
	// 		PrimaryKey: to.Ptr("############################################"),
	// 		SecondaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-8980.servicebus.windows-int.net/;SharedAccessKeyName=sdk-Authrules-8929;SharedAccessKey=############################################"),
	// 		SecondaryKey: to.Ptr("############################################"),
	// 	},
	// }
}
