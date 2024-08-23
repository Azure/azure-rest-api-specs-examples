const { HDInsightContainersManagementClient } = require("@azure/arm-hdinsightcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all libraries of HDInsight on AKS cluster.
 *
 * @summary Get all libraries of HDInsight on AKS cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/ListUserCustomClusterLibraries.json
 */
async function listUserCustomClusterLibraries() {
  const subscriptionId =
    process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "hiloResourceGroup";
  const clusterPoolName = "clusterPool";
  const clusterName = "cluster";
  const category = "custom";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightContainersManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clusterLibraries.list(
    resourceGroupName,
    clusterPoolName,
    clusterName,
    category,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
