const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a TagRule
 *
 * @summary Update a TagRule
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/TagRules_Update_MaximumSet_Gen.json
 */
async function tagRulesUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["DYNATRACE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["DYNATRACE_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "myMonitor";
  const ruleSetName = "default";
  const resource = {
    logRules: {
      filteringTags: [
        { name: "Environment", action: "Include", value: "Prod" },
        { name: "Environment", action: "Exclude", value: "Dev" },
      ],
      sendAadLogs: "Enabled",
      sendActivityLogs: "Enabled",
      sendSubscriptionLogs: "Enabled",
    },
    metricRules: {
      filteringTags: [{ name: "Environment", action: "Include", value: "Prod" }],
    },
  };
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
