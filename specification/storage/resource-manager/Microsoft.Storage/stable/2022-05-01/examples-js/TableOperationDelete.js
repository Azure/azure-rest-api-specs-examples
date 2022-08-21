const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the table with the specified table name, under the specified account if it exists.
 *
 * @summary Deletes the table with the specified table name, under the specified account if it exists.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/TableOperationDelete.json
 */
async function tableOperationDelete() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
  const accountName = "sto328";
  const tableName = "table6185";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.tableOperations.delete(resourceGroupName, accountName, tableName);
  console.log(result);
}

tableOperationDelete().catch(console.error);
