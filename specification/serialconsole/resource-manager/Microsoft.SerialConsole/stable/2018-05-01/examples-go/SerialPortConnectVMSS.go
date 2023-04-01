package armserialconsole_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/serialconsole/armserialconsole"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/SerialPortConnectVMSS.json
func ExampleSerialPortsClient_Connect_connectToAScaleSetInstanceSerialPort() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armserialconsole.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSerialPortsClient().Connect(ctx, "myResourceGroup", "Microsoft.Compute", "virtualMachineScaleSets", "myscaleset/virtualMachines/2", "0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SerialPortConnectResult = armserialconsole.SerialPortConnectResult{
	// 	ConnectionString: to.Ptr("wss://eastus.gateway.serialconsole.azure.com/n/connector/{containerid}/sessions/{sessionId}/client"),
	// }
}
