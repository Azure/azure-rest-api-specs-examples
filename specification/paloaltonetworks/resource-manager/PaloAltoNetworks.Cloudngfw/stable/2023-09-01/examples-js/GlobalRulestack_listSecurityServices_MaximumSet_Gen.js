const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the security services for rulestack
 *
 * @summary List the security services for rulestack
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/GlobalRulestack_listSecurityServices_MaximumSet_Gen.json
 */
async function globalRulestackListSecurityServicesMaximumSetGen() {
  const globalRulestackName = "praval";
  const skip = "a6a321";
  const top = 20;
  const typeParam = "globalRulestacks";
  const options = {
    skip,
    top,
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.listSecurityServices(
    globalRulestackName,
    typeParam,
    options
  );
  console.log(result);
}
