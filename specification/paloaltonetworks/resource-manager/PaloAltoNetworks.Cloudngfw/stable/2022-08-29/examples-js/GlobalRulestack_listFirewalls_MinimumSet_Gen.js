const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of Firewalls associated with Rulestack
 *
 * @summary List of Firewalls associated with Rulestack
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/GlobalRulestack_listFirewalls_MinimumSet_Gen.json
 */
async function globalRulestackListFirewallsMinimumSetGen() {
  const globalRulestackName = "praval";
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.listFirewalls(globalRulestackName);
  console.log(result);
}
