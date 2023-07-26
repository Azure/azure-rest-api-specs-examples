const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all scripts' execution history for the specified cluster.
 *
 * @summary Lists all scripts' execution history for the specified cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/GetScriptExecutionHistory.json
 */
async function getScriptExecutionHistoryList() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.scriptExecutionHistory.listByCluster(
    resourceGroupName,
    clusterName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
