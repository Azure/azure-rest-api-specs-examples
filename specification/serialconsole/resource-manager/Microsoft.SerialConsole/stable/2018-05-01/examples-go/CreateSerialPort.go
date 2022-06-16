package armserialconsole_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/serialconsole/armserialconsole"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/CreateSerialPort.json
func ExampleSerialPortsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armserialconsole.NewSerialPortsClient("00000000-00000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx,
		"myResourceGroup",
		"Microsoft.Compute",
		"virtualMachines",
		"myVM",
		"0",
		armserialconsole.SerialPort{
			Properties: &armserialconsole.SerialPortProperties{
				State: to.Ptr(armserialconsole.SerialPortStateEnabled),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
