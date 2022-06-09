```go
package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/disasterRecoveryConfigs/SBAliasFailOver.json
func ExampleDisasterRecoveryConfigsClient_FailOver() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewDisasterRecoveryConfigsClient("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.FailOver(ctx,
		"ardsouzatestRG",
		"sdk-Namespace-8860",
		"sdk-DisasterRecovery-3814",
		&armservicebus.DisasterRecoveryConfigsClientFailOverOptions{Parameters: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicebus%2Farmservicebus%2Fv2.0.0-beta.1/sdk/resourcemanager/servicebus/armservicebus/README.md) on how to add the SDK to your project and authenticate.
