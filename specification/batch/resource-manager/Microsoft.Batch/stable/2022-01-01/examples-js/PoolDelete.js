const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified pool.
 *
 * @summary Deletes the specified pool.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PoolDelete.json
 */
async function deletePool() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.poolOperations.beginDeleteAndWait(
    resourceGroupName,
    accountName,
    poolName
  );
  console.log(result);
}

deletePool().catch(console.error);
