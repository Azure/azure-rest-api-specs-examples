const { MicrosoftSerialConsoleClient } = require("@azure/arm-serialconsole");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to connect to serial port of the target resource
 *
 * @summary connect to serial port of the target resource
 * x-ms-original-file: 2024-07-01/SerialPortConnectVM.json
 */
async function connectToAVirtualMachineSerialPort() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new MicrosoftSerialConsoleClient(credential, subscriptionId);
  const result = await client.serialPorts.connect(
    "myResourceGroup",
    "Microsoft.Compute",
    "virtualMachines",
    "myVM",
    "0",
  );
  console.log(result);
}
