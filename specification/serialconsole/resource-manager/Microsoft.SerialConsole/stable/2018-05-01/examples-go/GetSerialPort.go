package armserialconsole_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/serialconsole/armserialconsole"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/GetSerialPort.json
func ExampleSerialPortsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armserialconsole.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSerialPortsClient().Get(ctx, "myResourceGroup", "Microsoft.Compute", "virtualMachines", "myVM", "0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SerialPort = armserialconsole.SerialPort{
	// 	Name: to.Ptr("0"),
	// 	Type: to.Ptr("Microsoft.SerialConsole/serialPorts"),
	// 	ID: to.Ptr("/subscriptions/00000000-00000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/providers/Microsoft.SerialConsole/serialPorts/0"),
	// 	Properties: &armserialconsole.SerialPortProperties{
	// 		State: to.Ptr(armserialconsole.SerialPortStateEnabled),
	// 	},
	// }
}
