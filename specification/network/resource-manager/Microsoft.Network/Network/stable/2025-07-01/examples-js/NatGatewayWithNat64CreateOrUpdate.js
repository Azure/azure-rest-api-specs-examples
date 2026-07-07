const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a nat gateway.
 *
 * @summary creates or updates a nat gateway.
 * x-ms-original-file: 2025-07-01/NatGatewayWithNat64CreateOrUpdate.json
 */
async function createNatGatewayWithNat64() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.natGateways.createOrUpdate("rg1", "test-natgateway", {
    location: "westus",
    nat64: "Enabled",
    publicIpAddresses: [
      {
        id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1",
      },
    ],
    publicIpPrefixes: [
      {
        id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/PublicIpPrefix1",
      },
    ],
    sku: { name: "Standard" },
  });
  console.log(result);
}
