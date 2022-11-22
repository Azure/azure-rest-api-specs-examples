const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the given detector for a given Batch account.
 *
 * @summary Gets information about the given detector for a given Batch account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/DetectorGet.json
 */
async function getDetector() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const detectorId = "poolsAndNodes";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.batchAccountOperations.getDetector(
    resourceGroupName,
    accountName,
    detectorId
  );
  console.log(result);
}

getDetector().catch(console.error);
