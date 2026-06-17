const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a PostRulesResource
 *
 * @summary delete a PostRulesResource
 * x-ms-original-file: 2026-05-11-preview/PostRules_Delete_MaximumSet_Gen.json
 */
async function postRulesDeleteMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  await client.postRules.delete("lrs1", "1");
}
