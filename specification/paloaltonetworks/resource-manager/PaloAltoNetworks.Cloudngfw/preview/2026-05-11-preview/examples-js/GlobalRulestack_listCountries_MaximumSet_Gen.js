const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list of countries for Rulestack
 *
 * @summary list of countries for Rulestack
 * x-ms-original-file: 2026-05-11-preview/GlobalRulestack_listCountries_MaximumSet_Gen.json
 */
async function globalRulestackListCountriesMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.listCountries("praval", { skip: "a6a321", top: 20 });
  console.log(result);
}
