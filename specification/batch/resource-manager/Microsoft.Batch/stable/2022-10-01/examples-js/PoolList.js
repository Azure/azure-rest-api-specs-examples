const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the pools in the specified account.
 *
 * @summary Lists all of the pools in the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolList.json
 */
async function listPool() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.poolOperations.listByBatchAccount(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPool().catch(console.error);
