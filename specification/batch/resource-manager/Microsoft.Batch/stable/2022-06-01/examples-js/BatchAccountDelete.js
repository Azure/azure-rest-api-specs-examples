const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified Batch account.
 *
 * @summary Deletes the specified Batch account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-06-01/examples/BatchAccountDelete.json
 */
async function batchAccountDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.batchAccountOperations.beginDeleteAndWait(
    resourceGroupName,
    accountName
  );
  console.log(result);
}

batchAccountDelete().catch(console.error);
