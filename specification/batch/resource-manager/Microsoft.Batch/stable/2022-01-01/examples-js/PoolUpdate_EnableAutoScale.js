const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the properties of an existing pool.
 *
 * @summary Updates the properties of an existing pool.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PoolUpdate_EnableAutoScale.json
 */
async function updatePoolEnableAutoscale() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const parameters = {
    scaleSettings: { autoScale: { formula: "$TargetDedicatedNodes=34" } },
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.poolOperations.update(
    resourceGroupName,
    accountName,
    poolName,
    parameters
  );
  console.log(result);
}

updatePoolEnableAutoscale().catch(console.error);
