const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves all nat rules for a particular virtual wan vpn gateway.
 *
 * @summary Retrieves all nat rules for a particular virtual wan vpn gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NatRuleList.json
 */
async function natRuleList() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "gateway1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.natRules.listByVpnGateway(resourceGroupName, gatewayName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

natRuleList().catch(console.error);
