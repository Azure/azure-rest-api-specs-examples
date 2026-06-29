const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new table with the specified table name, under the specified account.
 *
 * @summary creates a new table with the specified table name, under the specified account.
 * x-ms-original-file: 2026-04-01/TableOperationPatch.json
 */
async function tableOperationPatch() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.table.update("res3376", "sto328", "table6185");
  console.log(result);
}
