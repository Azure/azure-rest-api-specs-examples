const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified cluster.
 *
 * @summary gets the specified cluster.
 * x-ms-original-file: 2025-01-15-preview/GetLinuxHadoopCluster.json
 */
async function getHadoopOnLinuxCluster() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.get("rg1", "cluster1");
  console.log(result);
}
