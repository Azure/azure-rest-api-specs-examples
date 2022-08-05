const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the extension properties for the specified HDInsight cluster extension.
 *
 * @summary Gets the extension properties for the specified HDInsight cluster extension.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetExtension.json
 */
async function getAnExtension() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cluster1";
  const extensionName = "clustermonitoring";
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.extensions.get(resourceGroupName, clusterName, extensionName);
  console.log(result);
}

getAnExtension().catch(console.error);
