const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates the tag rules for a specific New Relic monitor resource, allowing you to modify the rules that control which Azure resources are monitored
 *
 * @summary updates the tag rules for a specific New Relic monitor resource, allowing you to modify the rules that control which Azure resources are monitored
 * x-ms-original-file: 2025-05-01-preview/TagRules_Update_MinimumSet_Gen.json
 */
async function tagRulesUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NewRelicObservability(credential, subscriptionId);
  const result = await client.tagRules.update(
    "rgopenapi",
    "ipxmlcbonyxtolzejcjshkmlron",
    "bxcantgzggsepbhqmedjqyrqeezmfb",
    {},
  );
  console.log(result);
}
