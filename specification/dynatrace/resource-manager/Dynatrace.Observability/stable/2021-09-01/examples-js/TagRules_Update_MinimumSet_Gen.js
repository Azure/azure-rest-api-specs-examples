const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a TagRule
 *
 * @summary Update a TagRule
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/TagRules_Update_MinimumSet_Gen.json
 */
async function tagRulesUpdateMinimumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const ruleSetName = "default";
  const resource = {};
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.tagRules.update(
    resourceGroupName,
    monitorName,
    ruleSetName,
    resource
  );
  console.log(result);
}

tagRulesUpdateMinimumSetGen().catch(console.error);
