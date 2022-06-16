const { MicrosoftSerialConsoleClient } = require("@azure/arm-serialconsole");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of Serial Console API operations.
 *
 * @summary Gets a list of Serial Console API operations.
 * x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/GetOperationsExample.json
 */
async function listAllSerialConsoleManagementOperationsSupportedBySerialConsoleRp() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSerialConsoleClient(credential, subscriptionId);
  const result = await client.listOperations();
  console.log(result);
}

listAllSerialConsoleManagementOperationsSupportedBySerialConsoleRp().catch(console.error);
