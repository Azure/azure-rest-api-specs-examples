const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates the Autoscale Configuration for HDInsight cluster.
 *
 * @summary updates the Autoscale Configuration for HDInsight cluster.
 * x-ms-original-file: 2025-01-15-preview/EnableOrUpdateAutoScaleWithLoadBasedConfiguration.json
 */
async function enableOrUpdateAutoscaleWithTheLoadBasedConfigurationForHDInsightCluster() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new HDInsightManagementClient(credential, subscriptionId);
  await client.clusters.updateAutoScaleConfiguration("rg1", "cluster1", "workernode", {
    autoscale: { capacity: { maxInstanceCount: 5, minInstanceCount: 3 } },
  });
}
