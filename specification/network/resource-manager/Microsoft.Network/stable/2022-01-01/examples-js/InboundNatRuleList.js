const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the inbound NAT rules in a load balancer.
 *
 * @summary Gets all the inbound NAT rules in a load balancer.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/InboundNatRuleList.json
 */
async function inboundNatRuleList() {
  const subscriptionId = "subid";
  const resourceGroupName = "testrg";
  const loadBalancerName = "lb1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.inboundNatRules.list(resourceGroupName, loadBalancerName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

inboundNatRuleList().catch(console.error);
