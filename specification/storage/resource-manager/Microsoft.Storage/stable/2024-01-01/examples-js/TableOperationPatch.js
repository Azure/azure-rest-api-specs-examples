const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a new table with the specified table name, under the specified account.
 *
 * @summary Creates a new table with the specified table name, under the specified account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/TableOperationPatch.json
 */
async function tableOperationPatch() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res3376";
  const accountName = "sto328";
  const tableName = "table6185";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.tableOperations.update(resourceGroupName, accountName, tableName);
  console.log(result);
}
