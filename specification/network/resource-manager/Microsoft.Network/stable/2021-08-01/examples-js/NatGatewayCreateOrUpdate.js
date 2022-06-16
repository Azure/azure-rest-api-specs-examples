const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a nat gateway.
 *
 * @summary Creates or updates a nat gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NatGatewayCreateOrUpdate.json
 */
async function createNatGateway() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const natGatewayName = "test-natgateway";
  const parameters = {
    location: "westus",
    publicIpAddresses: [
      {
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1",
      },
    ],
    publicIpPrefixes: [
      {
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/PublicIpPrefix1",
      },
    ],
    sku: { name: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.natGateways.beginCreateOrUpdateAndWait(
    resourceGroupName,
    natGatewayName,
    parameters
  );
  console.log(result);
}

createNatGateway().catch(console.error);
