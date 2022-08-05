const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the Autoscale Configuration for HDInsight cluster.
 *
 * @summary Updates the Autoscale Configuration for HDInsight cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/EnableOrUpdateAutoScaleWithLoadBasedConfiguration.json
 */
async function enableOrUpdateAutoscaleWithTheLoadBasedConfigurationForHdInsightCluster() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cluster1";
  const roleName = "workernode";
  const parameters = {
    autoscale: { capacity: { maxInstanceCount: 5, minInstanceCount: 3 } },
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginUpdateAutoScaleConfigurationAndWait(
    resourceGroupName,
    clusterName,
    roleName,
    parameters
  );
  console.log(result);
}

enableOrUpdateAutoscaleWithTheLoadBasedConfigurationForHdInsightCluster().catch(console.error);
