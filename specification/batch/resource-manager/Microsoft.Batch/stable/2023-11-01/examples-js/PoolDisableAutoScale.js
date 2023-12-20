const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Disables automatic scaling for a pool.
 *
 * @summary Disables automatic scaling for a pool.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/PoolDisableAutoScale.json
 */
async function disableAutoScale() {
  const subscriptionId = process.env["BATCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["BATCH_RESOURCE_GROUP"] || "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.poolOperations.disableAutoScale(
    resourceGroupName,
    accountName,
    poolName,
  );
  console.log(result);
}
