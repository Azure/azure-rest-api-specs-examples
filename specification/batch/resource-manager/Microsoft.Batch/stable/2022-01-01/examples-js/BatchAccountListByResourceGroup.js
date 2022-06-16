const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the Batch accounts associated with the specified resource group.
 *
 * @summary Gets information about the Batch accounts associated with the specified resource group.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/BatchAccountListByResourceGroup.json
 */
async function batchAccountListByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.batchAccountOperations.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

batchAccountListByResourceGroup().catch(console.error);
