const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the status of Azure Monitor on the HDInsight cluster.
 *
 * @summary Gets the status of Azure Monitor on the HDInsight cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetLinuxClusterAzureMonitorStatus.json
 */
async function enableClusterMonitoring() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cluster1";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.extensions.getAzureMonitorStatus(resourceGroupName, clusterName);
  console.log(result);
}

enableClusterMonitoring().catch(console.error);
