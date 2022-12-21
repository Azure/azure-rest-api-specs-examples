const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the Firewall Policies in a subscription.
 *
 * @summary Gets all the Firewall Policies in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/FirewallPolicyListBySubscription.json
 */
async function listAllFirewallPoliciesForAGivenSubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.firewallPolicies.listAll()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllFirewallPoliciesForAGivenSubscription().catch(console.error);
