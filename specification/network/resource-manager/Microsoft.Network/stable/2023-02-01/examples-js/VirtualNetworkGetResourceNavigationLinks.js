const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of resource navigation links for a subnet.
 *
 * @summary Gets a list of resource navigation links for a subnet.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/VirtualNetworkGetResourceNavigationLinks.json
 */
async function getResourceNavigationLinks() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualNetworkName = "vnet";
  const subnetName = "subnet";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.resourceNavigationLinks.list(
    resourceGroupName,
    virtualNetworkName,
    subnetName
  );
  console.log(result);
}
