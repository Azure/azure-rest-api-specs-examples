const { MicrosoftSerialConsoleClient } = require("@azure/arm-serialconsole");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to connect to serial port of the target resource
 *
 * @summary connect to serial port of the target resource
 * x-ms-original-file: 2024-07-01/SerialPortConnectVMSS.json
 */
async function connectToAScaleSetInstanceSerialPort() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new MicrosoftSerialConsoleClient(credential, subscriptionId);
  const result = await client.serialPorts.connect(
    "myResourceGroup",
    "Microsoft.Compute",
    "virtualMachineScaleSets",
    "myscaleset/virtualMachines/2",
    "0",
  );
  console.log(result);
}
