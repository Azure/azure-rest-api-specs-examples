const { MicrosoftSerialConsoleClient } = require("@azure/arm-serialconsole");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Disables the Serial Console service for all VMs and VM scale sets in the provided subscription
 *
 * @summary Disables the Serial Console service for all VMs and VM scale sets in the provided subscription
 * x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/DisableConsoleExamples.json
 */
async function disableSerialConsoleForASubscription() {
  const subscriptionId = "00000000-00000-0000-0000-000000000000";
  const defaultParam = "default";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSerialConsoleClient(credential, subscriptionId);
  const result = await client.disableConsole(defaultParam);
  console.log(result);
}

disableSerialConsoleForASubscription().catch(console.error);
