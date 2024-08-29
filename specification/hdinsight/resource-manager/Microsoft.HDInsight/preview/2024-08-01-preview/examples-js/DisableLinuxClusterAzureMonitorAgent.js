const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Disables the Azure Monitor Agent on the HDInsight cluster.
 *
 * @summary Disables the Azure Monitor Agent on the HDInsight cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/DisableLinuxClusterAzureMonitorAgent.json
 */
async function disableAzureMonitorAgent() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.extensions.beginDisableAzureMonitorAgentAndWait(
    resourceGroupName,
    clusterName,
  );
  console.log(result);
}
