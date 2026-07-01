const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete the connection type.
 *
 * @summary delete the connection type.
 * x-ms-original-file: 2024-10-23/deleteConnectionType.json
 */
async function deleteAnExistingConnectionType() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee";
  const client = new AutomationClient(credential, subscriptionId);
  await client.connectionTypeOperations.delete("rg", "myAutomationAccount22", "myCT");
}
