const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List the security services for rulestack
 *
 * @summary List the security services for rulestack
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2025-10-08/examples/GlobalRulestack_listSecurityServices_MinimumSet_Gen.json
 */
async function globalRulestackListSecurityServicesMinimumSetGen() {
  const globalRulestackName = "praval";
  const typeParam = "globalRulestacks";
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.listSecurityServices(globalRulestackName, typeParam);
  console.log(result);
}
