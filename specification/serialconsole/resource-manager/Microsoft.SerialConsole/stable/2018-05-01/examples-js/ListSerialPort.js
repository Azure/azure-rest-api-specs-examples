const { MicrosoftSerialConsoleClient } = require("@azure/arm-serialconsole");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the configured serial ports for a parent resource
 *
 * @summary Lists all of the configured serial ports for a parent resource
 * x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/ListSerialPort.json
 */
async function listSerialPortsForParentResources() {
  const subscriptionId = "00000000-00000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceProviderNamespace = "Microsoft.Compute";
  const parentResourceType = "virtualMachines";
  const parentResource = "myVM";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSerialConsoleClient(credential, subscriptionId);
  const result = await client.serialPorts.list(
    resourceGroupName,
    resourceProviderNamespace,
    parentResourceType,
    parentResource
  );
  console.log(result);
}

listSerialPortsForParentResources().catch(console.error);
