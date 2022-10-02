const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified load balancer load balancing rule.
 *
 * @summary Gets the specified load balancer load balancing rule.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/LoadBalancerLoadBalancingRuleGet.json
 */
async function loadBalancerLoadBalancingRuleGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "testrg";
  const loadBalancerName = "lb1";
  const loadBalancingRuleName = "rule1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.loadBalancerLoadBalancingRules.get(
    resourceGroupName,
    loadBalancerName,
    loadBalancingRuleName
  );
  console.log(result);
}

loadBalancerLoadBalancingRuleGet().catch(console.error);
