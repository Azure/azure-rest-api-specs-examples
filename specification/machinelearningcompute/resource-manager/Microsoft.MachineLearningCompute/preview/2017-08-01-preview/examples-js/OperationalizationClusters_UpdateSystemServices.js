const { MachineLearningComputeManagementClient } = require("@azure/arm-machinelearningcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates system services in a cluster.
 *
 * @summary Updates system services in a cluster.
 * x-ms-original-file: specification/machinelearningcompute/resource-manager/Microsoft.MachineLearningCompute/preview/2017-08-01-preview/examples/OperationalizationClusters_UpdateSystemServices.json
 */
async function updateSystemServicesInAnOperationalizationCluster() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const clusterName = "myCluster";
  const credential = new DefaultAzureCredential();
  const client = new MachineLearningComputeManagementClient(credential, subscriptionId);
  const result = await client.operationalizationClusters.beginUpdateSystemServicesAndWait(
    resourceGroupName,
    clusterName
  );
  console.log(result);
}

updateSystemServicesInAnOperationalizationCluster().catch(console.error);
