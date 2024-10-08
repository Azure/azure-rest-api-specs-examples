const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restarts the specified HDInsight cluster hosts.
 *
 * @summary Restarts the specified HDInsight cluster hosts.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/RestartVirtualMachinesOperation.json
 */
async function restartsTheSpecifiedHdInsightClusterHosts() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const hosts = ["gateway1", "gateway3"];
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginRestartHostsAndWait(
    resourceGroupName,
    clusterName,
    hosts,
  );
  console.log(result);
}
