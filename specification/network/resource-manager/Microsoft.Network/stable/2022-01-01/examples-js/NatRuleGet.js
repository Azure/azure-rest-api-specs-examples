const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the details of a nat ruleGet.
 *
 * @summary Retrieves the details of a nat ruleGet.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NatRuleGet.json
 */
async function natRuleGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "gateway1";
  const natRuleName = "natRule1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.natRules.get(resourceGroupName, gatewayName, natRuleName);
  console.log(result);
}

natRuleGet().catch(console.error);
