package armserialconsole_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/serialconsole/armserialconsole/v2"
)

// Generated from example definition: 2024-07-01/CreateSerialPort.json
func ExampleSerialPortsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armserialconsole.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSerialPortsClient().Create(ctx, "myResourceGroup", "Microsoft.Compute", "virtualMachines", "myVM", "0", armserialconsole.SerialPort{
		Properties: &armserialconsole.SerialPortProperties{
			State: to.Ptr(armserialconsole.SerialPortStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
