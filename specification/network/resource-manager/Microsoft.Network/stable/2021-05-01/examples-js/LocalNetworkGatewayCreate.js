const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a local network gateway in the specified resource group.
 *
 * @summary Creates or updates a local network gateway in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LocalNetworkGatewayCreate.json
 */
async function createLocalNetworkGateway() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const localNetworkGatewayName = "localgw";
  const parameters = {
    fqdn: "site1.contoso.com",
    gatewayIpAddress: "11.12.13.14",
    localNetworkAddressSpace: { addressPrefixes: ["10.1.0.0/16"] },
    location: "Central US",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.localNetworkGateways.beginCreateOrUpdateAndWait(
    resourceGroupName,
    localNetworkGatewayName,
    parameters
  );
  console.log(result);
}

createLocalNetworkGateway().catch(console.error);
