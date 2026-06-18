const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a PreRulesResource
 *
 * @summary get a PreRulesResource
 * x-ms-original-file: 2026-05-11-preview/PreRules_Get_MinimumSet_Gen.json
 */
async function preRulesGetMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.preRules.get("lrs1", "1");
  console.log(result);
}
