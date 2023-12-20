const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the properties of an existing pool.
 *
 * @summary Updates the properties of an existing pool.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/PoolUpdate_ResizePool.json
 */
async function updatePoolResizePool() {
  const subscriptionId = process.env["BATCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["BATCH_RESOURCE_GROUP"] || "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const parameters = {
    scaleSettings: {
      fixedScale: {
        nodeDeallocationOption: "TaskCompletion",
        resizeTimeout: "PT8M",
        targetDedicatedNodes: 5,
        targetLowPriorityNodes: 0,
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.poolOperations.update(
    resourceGroupName,
    accountName,
    poolName,
    parameters,
  );
  console.log(result);
}
