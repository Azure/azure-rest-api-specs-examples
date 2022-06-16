const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of service association links for a subnet.
 *
 * @summary Gets a list of service association links for a subnet.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGetServiceAssociationLinks.json
 */
async function getServiceAssociationLinks() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkName = "vnet";
  const subnetName = "subnet";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.serviceAssociationLinks.list(
    resourceGroupName,
    virtualNetworkName,
    subnetName
  );
  console.log(result);
}

getServiceAssociationLinks().catch(console.error);
