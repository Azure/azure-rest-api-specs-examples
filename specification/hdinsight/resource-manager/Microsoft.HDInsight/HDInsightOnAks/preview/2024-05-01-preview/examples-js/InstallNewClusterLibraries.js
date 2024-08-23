const { HDInsightContainersManagementClient } = require("@azure/arm-hdinsightcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Library management operations on HDInsight on AKS cluster.
 *
 * @summary Library management operations on HDInsight on AKS cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/InstallNewClusterLibraries.json
 */
async function installNewClusterLibraries() {
  const subscriptionId =
    process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "hiloResourceGroup";
  const clusterPoolName = "clusterPool";
  const clusterName = "cluster";
  const operation = {
    properties: {
      action: "Install",
      libraries: [
        {
          properties: {
            name: "requests",
            type: "pypi",
            remarks: "PyPi packages.",
            version: "2.31.0",
          },
        },
        {
          properties: {
            name: "flink-connector-kafka",
            type: "maven",
            groupId: "org.apache.flink",
            remarks: "Maven packages.",
            version: "3.0.2-1.18",
          },
        },
      ],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightContainersManagementClient(credential, subscriptionId);
  const result = await client.clusterLibraries.beginManageLibrariesAndWait(
    resourceGroupName,
    clusterPoolName,
    clusterName,
    operation,
  );
  console.log(result);
}
