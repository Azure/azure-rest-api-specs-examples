package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/HybridConnection/RelayHybridConnectionAuthorizationRuleRegenerateKey.json
func ExampleHybridConnectionsClient_RegenerateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHybridConnectionsClient().RegenerateKeys(ctx, "resourcegroup", "example-RelayNamespace-01", "example-Relay-Hybrid-01", "example-RelayAuthRules-01", armrelay.RegenerateAccessKeyParameters{
		KeyType: to.Ptr(armrelay.KeyTypePrimaryKey),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessKeys = armrelay.AccessKeys{
	// 	KeyName: to.Ptr("example-RelayAuthRules-01"),
	// 	PrimaryConnectionString: to.Ptr("Endpoint=sb://example-Relaynamespace-01.servicebus.windows.net/;SharedAccessKeyName=example-RelayAuthRules-01;SharedAccessKey=############################################"),
	// 	PrimaryKey: to.Ptr("############################################"),
	// 	SecondaryConnectionString: to.Ptr("Endpoint=sb://example-Relaynamespace-01.servicebus.windows.net/;SharedAccessKeyName=example-RelayAuthRules-01;SharedAccessKey=############################################"),
	// 	SecondaryKey: to.Ptr("############################################"),
	// }
}
