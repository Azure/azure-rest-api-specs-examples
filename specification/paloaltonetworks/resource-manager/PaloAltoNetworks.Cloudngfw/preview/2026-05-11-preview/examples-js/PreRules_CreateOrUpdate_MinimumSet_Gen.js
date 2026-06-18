const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a PreRulesResource
 *
 * @summary create a PreRulesResource
 * x-ms-original-file: 2026-05-11-preview/PreRules_CreateOrUpdate_MinimumSet_Gen.json
 */
async function preRulesCreateOrUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.preRules.createOrUpdate("lrs1", "1", { ruleName: "preRule1" });
  console.log(result);
}
