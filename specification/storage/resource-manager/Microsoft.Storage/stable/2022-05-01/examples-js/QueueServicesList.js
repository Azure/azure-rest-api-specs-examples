const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all queue services for the storage account
 *
 * @summary List all queue services for the storage account
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/QueueServicesList.json
 */
async function queueServicesList() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9290";
  const accountName = "sto1590";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.queueServices.list(resourceGroupName, accountName);
  console.log(result);
}

queueServicesList().catch(console.error);
