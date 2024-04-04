const { HDInsightContainersManagementClient } = require("@azure/arm-hdinsightcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a cluster pool.
 *
 * @summary Creates or updates a cluster pool.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-11-01-preview/examples/CreateClusterPoolWithUDRAks.json
 */
async function clusterPoolPutWithUdrAks() {
  const subscriptionId =
    process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "10e32bab-26da-4cc4-a441-52b318f824e6";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "hiloResourcegroup";
  const clusterPoolName = "clusterpool1";
  const clusterPool = {
    location: "West US 2",
    properties: {
      clusterPoolProfile: { clusterPoolVersion: "1.2" },
      computeProfile: { vmSize: "Standard_D3_v2" },
      networkProfile: {
        outboundType: "userDefinedRouting",
        subnetId:
          "/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightContainersManagementClient(credential, subscriptionId);
  const result = await client.clusterPools.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterPoolName,
    clusterPool,
  );
  console.log(result);
}
