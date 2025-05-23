const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a nat gateway.
 *
 * @summary Creates or updates a nat gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NatGatewayCreateOrUpdateStandardV2Sku.json
 */
async function createNatGatewayWithStandardV2Sku() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
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
    sku: { name: "StandardV2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.natGateways.beginCreateOrUpdateAndWait(
    resourceGroupName,
    natGatewayName,
    parameters,
  );
  console.log(result);
}
