const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update tags associated with the additional details related to the Hybrid AKS provisioned cluster.
 *
 * @summary Update tags associated with the additional details related to the Hybrid AKS provisioned cluster.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/HybridAksClusters_Patch.json
 */
async function patchHybridAksProvisionedClusterData() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const hybridAksClusterName = "hybridAksClusterName";
  const hybridAksClusterUpdateParameters = {
    tags: { key1: "myvalue1", key2: "myvalue2" },
  };
  const options = {
    hybridAksClusterUpdateParameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.hybridAksClusters.update(
    resourceGroupName,
    hybridAksClusterName,
    options
  );
  console.log(result);
}
