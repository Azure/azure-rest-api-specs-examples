const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

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
