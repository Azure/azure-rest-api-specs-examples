const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all effective security admin rules applied on a virtual network.
 *
 * @summary List all effective security admin rules applied on a virtual network.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerEffectiveSecurityAdminRulesList.json
 */
async function listEffectiveSecurityAdminRules() {
  const subscriptionId = "subscriptionA";
  const resourceGroupName = "myResourceGroup";
  const virtualNetworkName = "testVirtualNetwork";
  const parameters = { skipToken: "FakeSkipTokenCode" };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.listNetworkManagerEffectiveSecurityAdminRules(
    resourceGroupName,
    virtualNetworkName,
    parameters
  );
  console.log(result);
}

listEffectiveSecurityAdminRules().catch(console.error);
