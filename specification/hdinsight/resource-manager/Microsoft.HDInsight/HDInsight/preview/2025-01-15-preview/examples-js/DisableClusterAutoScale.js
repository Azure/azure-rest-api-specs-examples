const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates the Autoscale Configuration for HDInsight cluster.
 *
 * @summary updates the Autoscale Configuration for HDInsight cluster.
 * x-ms-original-file: 2025-01-15-preview/DisableClusterAutoScale.json
 */
async function disableAutoscaleForTheHDInsightCluster() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new HDInsightManagementClient(credential, subscriptionId);
  await client.clusters.updateAutoScaleConfiguration("rg1", "cluster1", "workernode", {});
}
