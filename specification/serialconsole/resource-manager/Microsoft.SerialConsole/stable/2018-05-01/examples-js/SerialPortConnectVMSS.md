Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-serialconsole_2.0.1/sdk/serialconsole/arm-serialconsole/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MicrosoftSerialConsoleClient } = require("@azure/arm-serialconsole");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Connect to serial port of the target resource
 *
 * @summary Connect to serial port of the target resource
 * x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/SerialPortConnectVMSS.json
 */
async function connectToAScaleSetInstanceSerialPort() {
  const subscriptionId = "00000000-00000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceProviderNamespace = "Microsoft.Compute";
  const parentResourceType = "virtualMachineScaleSets";
  const parentResource = "myscaleset/virtualMachines/2";
  const serialPort = "0";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSerialConsoleClient(credential, subscriptionId);
  const result = await client.serialPorts.connect(
    resourceGroupName,
    resourceProviderNamespace,
    parentResourceType,
    parentResource,
    serialPort
  );
  console.log(result);
}

connectToAScaleSetInstanceSerialPort().catch(console.error);
```
