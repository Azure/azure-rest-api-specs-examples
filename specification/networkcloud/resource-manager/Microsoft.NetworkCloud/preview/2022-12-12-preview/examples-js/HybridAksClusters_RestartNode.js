const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restart a targeted node of a Hybrid AKS cluster.
 *
 * @summary Restart a targeted node of a Hybrid AKS cluster.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/HybridAksClusters_RestartNode.json
 */
async function restartAHybridAksClusterNode() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const hybridAksClusterName = "hybridAksClusterName";
  const hybridAksClusterRestartNodeParameters = {
    nodeName: "nodeName",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.hybridAksClusters.beginRestartNodeAndWait(
    resourceGroupName,
    hybridAksClusterName,
    hybridAksClusterRestartNodeParameters
  );
  console.log(result);
}
