const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a TagRule
 *
 * @summary Delete a TagRule
 * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/preview/2022-07-01-preview/examples/TagRules_Delete_MinimumSet_Gen.json
 */
async function tagRulesDeleteMinimumSetGen() {
  const subscriptionId = process.env["NEWRELICOBSERVABILITY_SUBSCRIPTION_ID"] || "ddqonpqwjr";
  const resourceGroupName = process.env["NEWRELICOBSERVABILITY_RESOURCE_GROUP"] || "rgopenapi";
  const monitorName = "ipxmlcbonyxtolzejcjshkmlron";
  const ruleSetName = "bxcantgzggsepbhqmedjqyrqeezmfb";
  const credential = new DefaultAzureCredential();
  const client = new NewRelicObservability(credential, subscriptionId);
  const result = await client.tagRules.beginDeleteAndWait(
    resourceGroupName,
    monitorName,
    ruleSetName
  );
  console.log(result);
}
