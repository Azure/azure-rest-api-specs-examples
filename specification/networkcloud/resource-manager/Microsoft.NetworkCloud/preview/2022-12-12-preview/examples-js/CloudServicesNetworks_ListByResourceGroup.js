const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of cloud services networks in the provided resource group.
 *
 * @summary Get a list of cloud services networks in the provided resource group.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/CloudServicesNetworks_ListByResourceGroup.json
 */
async function listCloudServicesNetworksForResourceGroup() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.cloudServicesNetworks.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
