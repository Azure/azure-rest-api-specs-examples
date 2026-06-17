const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a PostRulesResource
 *
 * @summary create a PostRulesResource
 * x-ms-original-file: 2026-05-11-preview/PostRules_CreateOrUpdate_MinimumSet_Gen.json
 */
async function postRulesCreateOrUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.postRules.createOrUpdate("lrs1", "1", { ruleName: "postRule1" });
  console.log(result);
}
