const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified load balancer inbound NAT rule.
 *
 * @summary Gets the specified load balancer inbound NAT rule.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/InboundNatRuleGet.json
 */
async function inboundNatRuleGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "testrg";
  const loadBalancerName = "lb1";
  const inboundNatRuleName = "natRule1.1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.inboundNatRules.get(
    resourceGroupName,
    loadBalancerName,
    inboundNatRuleName
  );
  console.log(result);
}

inboundNatRuleGet().catch(console.error);
