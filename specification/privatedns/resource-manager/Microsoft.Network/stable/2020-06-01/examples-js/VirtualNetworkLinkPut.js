const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a virtual network link to the specified Private DNS zone.
 *
 * @summary Creates or updates a virtual network link to the specified Private DNS zone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/VirtualNetworkLinkPut.json
 */
async function putPrivateDnsZoneVirtualNetworkLink() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const virtualNetworkLinkName = "virtualNetworkLink1";
  const parameters = {
    location: "Global",
    registrationEnabled: false,
    tags: { key1: "value1" },
    virtualNetwork: {
      id: "/subscriptions/virtualNetworkSubscriptionId/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/virtualNetworkName",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkLinks.beginCreateOrUpdateAndWait(
    resourceGroupName,
    privateZoneName,
    virtualNetworkLinkName,
    parameters
  );
  console.log(result);
}

putPrivateDnsZoneVirtualNetworkLink().catch(console.error);
