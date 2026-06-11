const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new set of tag rules for a specific New Relic monitor resource, determining which Azure resources are monitored based on their tags
 *
 * @summary creates a new set of tag rules for a specific New Relic monitor resource, determining which Azure resources are monitored based on their tags
 * x-ms-original-file: 2025-05-01-preview/TagRules_CreateOrUpdate_MinimumSet_Gen.json
 */
async function tagRulesCreateOrUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NewRelicObservability(credential, subscriptionId);
  const result = await client.tagRules.createOrUpdate(
    "rgopenapi",
    "ipxmlcbonyxtolzejcjshkmlron",
    "bxcantgzggsepbhqmedjqyrqeezmfb",
    {},
  );
  console.log(result);
}
