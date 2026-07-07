const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the specified Azure Firewall.
 *
 * @summary creates or updates the specified Azure Firewall.
 * x-ms-original-file: 2025-07-01/AzureFirewallPutWithAfcConfiguration.json
 */
async function createAzureFirewallWithAFCControlPlane() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.azureFirewalls.createOrUpdate(
    "rg1",
    "azurefirewall",
    {
      location: "West US",
      ipConfigurations: [
        {
          name: "azureFirewallIpConfiguration",
          publicIPAddress: {
            id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName",
          },
          subnet: {
            id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/AzureFirewallSubnet",
          },
        },
      ],
      sku: { name: "AZFW_VNet", tier: "Standard" },
      threatIntelMode: "Alert",
      tags: { key1: "value1" },
      zones: [],
    },
    { createAfcControlPlane: true },
  );
  console.log(result);
}
