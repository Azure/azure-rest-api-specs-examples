const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates the properties of an existing pool.
 *
 * @summary updates the properties of an existing pool.
 * x-ms-original-file: 2025-06-01/PoolUpdate_RemoveStartTask.json
 */
async function updatePoolRemoveStartTask() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789012";
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.pool.update(
    "default-azurebatch-japaneast",
    "sampleacct",
    "testpool",
    { startTask: {} },
  );
  console.log(result);
}
