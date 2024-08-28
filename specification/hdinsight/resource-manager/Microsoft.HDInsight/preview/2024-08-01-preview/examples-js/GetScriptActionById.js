const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the script execution detail for the given script execution ID.
 *
 * @summary Gets the script execution detail for the given script execution ID.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetScriptActionById.json
 */
async function getScriptExecutionHistoryByScriptId() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const scriptExecutionId = "391145124054712";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.scriptActions.getExecutionDetail(
    resourceGroupName,
    clusterName,
    scriptExecutionId,
  );
  console.log(result);
}
