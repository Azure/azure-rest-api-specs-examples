const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a TagRule
 *
 * @summary get a TagRule
 * x-ms-original-file: 2024-04-24/TagRules_Get_MaximumSet_Gen.json
 */
async function tagRulesGetMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.tagRules.get("myResourceGroup", "myMonitor", "default");
  console.log(result);
}
