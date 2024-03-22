const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List TagRule resources by NewRelicMonitorResource
 *
 * @summary List TagRule resources by NewRelicMonitorResource
 * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/TagRules_ListByNewRelicMonitorResource_MinimumSet_Gen.json
 */
async function tagRulesListByNewRelicMonitorResourceMinimumSetGen() {
  const subscriptionId =
    process.env["NEWRELICOBSERVABILITY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["NEWRELICOBSERVABILITY_RESOURCE_GROUP"] || "rgopenapi";
  const monitorName = "ipxmlcbonyxtolzejcjshkmlron";
  const credential = new DefaultAzureCredential();
  const client = new NewRelicObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.tagRules.listByNewRelicMonitorResource(
    resourceGroupName,
    monitorName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
