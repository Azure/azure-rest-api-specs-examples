const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a nat rule.
 *
 * @summary Deletes a nat rule.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/VirtualNetworkGatewayNatRuleDelete.json
 */
async function virtualNetworkGatewayNatRuleDelete() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualNetworkGatewayName = "gateway1";
  const natRuleName = "natRule1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGatewayNatRules.beginDeleteAndWait(
    resourceGroupName,
    virtualNetworkGatewayName,
    natRuleName
  );
  console.log(result);
}
