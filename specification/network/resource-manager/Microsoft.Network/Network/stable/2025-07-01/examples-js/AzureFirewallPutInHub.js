const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the specified Azure Firewall.
 *
 * @summary creates or updates the specified Azure Firewall.
 * x-ms-original-file: 2025-07-01/AzureFirewallPutInHub.json
 */
async function createAzureFirewallInVirtualHub() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.azureFirewalls.createOrUpdate("rg1", "azurefirewall", {
    location: "West US",
    firewallPolicy: {
      id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/policy1",
    },
    hubIPAddresses: { publicIPs: { addresses: [], count: 1 } },
    sku: { name: "AZFW_Hub", tier: "Standard" },
    threatIntelMode: "Alert",
    virtualHub: {
      id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1",
    },
    tags: { key1: "value1" },
    zones: [],
  });
  console.log(result);
}
