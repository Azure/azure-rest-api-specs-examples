const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves all nat rules for a particular virtual network gateway.
 *
 * @summary Retrieves all nat rules for a particular virtual network gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/VirtualNetworkGatewayNatRuleList.json
 */
async function virtualNetworkGatewayNatRuleList() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualNetworkGatewayName = "gateway1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualNetworkGatewayNatRules.listByVirtualNetworkGateway(
    resourceGroupName,
    virtualNetworkGatewayName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
