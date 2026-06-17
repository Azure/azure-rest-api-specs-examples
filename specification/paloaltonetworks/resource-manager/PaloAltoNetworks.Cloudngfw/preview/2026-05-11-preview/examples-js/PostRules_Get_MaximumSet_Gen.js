const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a PostRulesResource
 *
 * @summary get a PostRulesResource
 * x-ms-original-file: 2026-05-11-preview/PostRules_Get_MaximumSet_Gen.json
 */
async function postRulesGetMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.postRules.get("lrs1", "1");
  console.log(result);
}
