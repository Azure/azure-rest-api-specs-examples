const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the status of an asynchronous operation for the Azure HPC Cache
 *
 * @summary Gets the status of an asynchronous operation for the Azure HPC Cache
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/AscOperations_Get.json
 */
async function ascOperationsGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const location = "westus";
  const operationId = "testoperationid";
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const result = await client.ascOperations.get(location, operationId);
  console.log(result);
}

ascOperationsGet().catch(console.error);
