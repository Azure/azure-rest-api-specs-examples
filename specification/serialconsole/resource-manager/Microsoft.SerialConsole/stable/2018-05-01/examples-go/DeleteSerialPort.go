package armserialconsole_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/serialconsole/armserialconsole"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/DeleteSerialPort.json
func ExampleSerialPortsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armserialconsole.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSerialPortsClient().Delete(ctx, "myResourceGroup", "Microsoft.Compute", "virtualMachines", "myVM", "0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
