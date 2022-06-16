const { MicrosoftSerialConsoleClient } = require("@azure/arm-serialconsole");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a serial port
 *
 * @summary Creates or updates a serial port
 * x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/CreateSerialPort.json
 */
async function createANewSerialPortResource() {
  const subscriptionId = "00000000-00000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceProviderNamespace = "Microsoft.Compute";
  const parentResourceType = "virtualMachines";
  const parentResource = "myVM";
  const serialPort = "0";
  const parameters = { state: "enabled" };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSerialConsoleClient(credential, subscriptionId);
  const result = await client.serialPorts.create(
    resourceGroupName,
    resourceProviderNamespace,
    parentResourceType,
    parentResource,
    serialPort,
    parameters
  );
  console.log(result);
}

createANewSerialPortResource().catch(console.error);
