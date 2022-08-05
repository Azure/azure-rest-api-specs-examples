const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch HDInsight cluster with the specified parameters.
 *
 * @summary Patch HDInsight cluster with the specified parameters.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/PatchLinuxHadoopCluster.json
 */
async function patchHdInsightLinuxClusters() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cluster1";
  const parameters = {
    tags: { key1: "val1", key2: "val2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.update(resourceGroupName, clusterName, parameters);
  console.log(result);
}

patchHdInsightLinuxClusters().catch(console.error);
