package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: 2025-05-01-preview/disasterRecoveryConfigs/SBAliasAuthorizationRuleListKey.json
func ExampleDisasterRecoveryConfigsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDisasterRecoveryConfigsClient().ListKeys(ctx, "exampleResourceGroup", "sdk-Namespace-2702", "sdk-DisasterRecovery-4047", "sdk-Authrules-1746", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicebus.DisasterRecoveryConfigsClientListKeysResponse{
	// 	AccessKeys: armservicebus.AccessKeys{
	// 		AliasPrimaryConnectionString: to.Ptr("Endpoint=sb://sdk-disasterrecovery-4047.servicebus.windows-int.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=############################################"),
	// 		AliasSecondaryConnectionString: to.Ptr("Endpoint=sb://sdk-disasterrecovery-4047.servicebus.windows-int.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=############################################"),
	// 		KeyName: to.Ptr("sdk-Authrules-1746"),
	// 		PrimaryKey: to.Ptr("############################################"),
	// 		SecondaryKey: to.Ptr("############################################"),
	// 	},
	// }
}
