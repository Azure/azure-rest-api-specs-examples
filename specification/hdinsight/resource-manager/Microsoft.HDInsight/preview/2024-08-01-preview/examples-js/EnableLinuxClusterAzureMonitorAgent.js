const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enables the Azure Monitor Agent on the HDInsight cluster.
 *
 * @summary Enables the Azure Monitor Agent on the HDInsight cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/EnableLinuxClusterAzureMonitorAgent.json
 */
async function enableAzureMonitoringAgent() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const parameters = {
    primaryKey: "**********",
    workspaceId: "a2090ead-8c9f-4fba-b70e-533e3e003163",
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.extensions.beginEnableAzureMonitorAgentAndWait(
    resourceGroupName,
    clusterName,
    parameters,
  );
  console.log(result);
}
