const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of countries for Rulestack
 *
 * @summary List of countries for Rulestack
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/GlobalRulestack_listCountries_MaximumSet_Gen.json
 */
async function globalRulestackListCountriesMaximumSetGen() {
  const globalRulestackName = "praval";
  const skip = "a6a321";
  const top = 20;
  const options = { skip, top };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.listCountries(globalRulestackName, options);
  console.log(result);
}
