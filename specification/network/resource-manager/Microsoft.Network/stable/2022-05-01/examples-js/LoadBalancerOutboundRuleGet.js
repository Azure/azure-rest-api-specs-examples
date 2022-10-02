const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified load balancer outbound rule.
 *
 * @summary Gets the specified load balancer outbound rule.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/LoadBalancerOutboundRuleGet.json
 */
async function loadBalancerOutboundRuleGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "testrg";
  const loadBalancerName = "lb1";
  const outboundRuleName = "rule1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.loadBalancerOutboundRules.get(
    resourceGroupName,
    loadBalancerName,
    outboundRuleName
  );
  console.log(result);
}

loadBalancerOutboundRuleGet().catch(console.error);
