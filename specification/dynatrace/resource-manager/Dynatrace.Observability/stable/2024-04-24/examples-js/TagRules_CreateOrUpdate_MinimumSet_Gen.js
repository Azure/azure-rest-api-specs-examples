const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a TagRule
 *
 * @summary create a TagRule
 * x-ms-original-file: 2024-04-24/TagRules_CreateOrUpdate_MinimumSet_Gen.json
 */
async function tagRulesCreateOrUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.tagRules.createOrUpdate(
    "myResourceGroup",
    "myMonitor",
    "default",
    {},
  );
  console.log(result);
}
