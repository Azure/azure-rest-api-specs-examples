const { MicrosoftSerialConsoleClient } = require("@azure/arm-serialconsole");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets whether or not Serial Console is disabled for a given subscription
 *
 * @summary Gets whether or not Serial Console is disabled for a given subscription
 * x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/SerialConsoleStatus.json
 */
async function getTheSerialConsoleDisabledStatusForASubscription() {
  const subscriptionId = "00000000-00000-0000-0000-000000000000";
  const defaultParam = "default";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSerialConsoleClient(credential, subscriptionId);
  const result = await client.getConsoleStatus(defaultParam);
  console.log(result);
}

getTheSerialConsoleDisabledStatusForASubscription().catch(console.error);
