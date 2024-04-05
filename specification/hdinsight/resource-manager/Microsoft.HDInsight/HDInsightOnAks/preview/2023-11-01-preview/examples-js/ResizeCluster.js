const { HDInsightContainersManagementClient } = require("@azure/arm-hdinsightcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Resize an existing Cluster.
 *
 * @summary Resize an existing Cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-11-01-preview/examples/ResizeCluster.json
 */
async function hdInsightClusterResize() {
  const subscriptionId =
    process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "10e32bab-26da-4cc4-a441-52b318f824e6";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "hiloResourcegroup";
  const clusterPoolName = "clusterpool1";
  const clusterName = "cluster1";
  const clusterResizeRequest = {
    location: "West US 2",
    properties: { targetWorkerNodeCount: 5 },
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightContainersManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginResizeAndWait(
    resourceGroupName,
    clusterPoolName,
    clusterName,
    clusterResizeRequest,
  );
  console.log(result);
}
