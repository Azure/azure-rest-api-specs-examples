const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the inbound NAT rules in a load balancer.
 *
 * @summary Gets all the inbound NAT rules in a load balancer.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/InboundNatRuleList.json
 */
async function inboundNatRuleList() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "testrg";
  const loadBalancerName = "lb1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.inboundNatRules.list(resourceGroupName, loadBalancerName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
