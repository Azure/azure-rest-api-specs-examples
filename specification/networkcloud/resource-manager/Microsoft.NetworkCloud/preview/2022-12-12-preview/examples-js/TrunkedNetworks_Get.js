const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get properties of the provided trunked network.
 *
 * @summary Get properties of the provided trunked network.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/TrunkedNetworks_Get.json
 */
async function getTrunkedNetwork() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const trunkedNetworkName = "trunkedNetworkName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.trunkedNetworks.get(resourceGroupName, trunkedNetworkName);
  console.log(result);
}
