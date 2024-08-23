const { HDInsightContainersManagementClient } = require("@azure/arm-hdinsightcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Upgrade a cluster pool.
 *
 * @summary Upgrade a cluster pool.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/UpgradeAKSPatchVersionForClusterPool.json
 */
async function clusterPoolsUpgradeAksPatchVersion() {
  const subscriptionId =
    process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "10e32bab-26da-4cc4-a441-52b318f824e6";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "hiloResourcegroup";
  const clusterPoolName = "clusterpool1";
  const clusterPoolUpgradeRequest = {
    properties: {
      upgradeAllClusterNodes: false,
      upgradeClusterPool: true,
      upgradeType: "AKSPatchUpgrade",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightContainersManagementClient(credential, subscriptionId);
  const result = await client.clusterPools.beginUpgradeAndWait(
    resourceGroupName,
    clusterPoolName,
    clusterPoolUpgradeRequest,
  );
  console.log(result);
}
