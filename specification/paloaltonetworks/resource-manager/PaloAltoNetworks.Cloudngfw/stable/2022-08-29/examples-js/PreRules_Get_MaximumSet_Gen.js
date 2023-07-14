const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a PreRulesResource
 *
 * @summary Get a PreRulesResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/PreRules_Get_MaximumSet_Gen.json
 */
async function preRulesGetMaximumSetGen() {
  const globalRulestackName = "lrs1";
  const priority = "1";
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.preRules.get(globalRulestackName, priority);
  console.log(result);
}
