const { MachineLearningComputeManagementClient } = require("@azure/arm-machinelearningcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the clusters in the specified resource group.
 *
 * @summary Gets the clusters in the specified resource group.
 * x-ms-original-file: specification/machinelearningcompute/resource-manager/Microsoft.MachineLearningCompute/preview/2017-08-01-preview/examples/OperationalizationClusters_ListByResourceGroup.json
 */
async function listOperationalizationClustersByResourceGroup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new MachineLearningComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operationalizationClusters.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listOperationalizationClustersByResourceGroup().catch(console.error);
