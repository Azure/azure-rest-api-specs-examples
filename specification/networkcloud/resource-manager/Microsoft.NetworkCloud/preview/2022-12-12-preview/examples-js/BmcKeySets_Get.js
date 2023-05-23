const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get baseboard management controller key set of the provided cluster.
 *
 * @summary Get baseboard management controller key set of the provided cluster.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BmcKeySets_Get.json
 */
async function getBaseboardManagementControllerKeySetOfCluster() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const clusterName = "clusterName";
  const bmcKeySetName = "bmcKeySetName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.bmcKeySets.get(resourceGroupName, clusterName, bmcKeySetName);
  console.log(result);
}
