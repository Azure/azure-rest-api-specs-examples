const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List TagRule resources by NewRelicMonitorResource
 *
 * @summary List TagRule resources by NewRelicMonitorResource
 * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/preview/2022-07-01-preview/examples/TagRules_ListByNewRelicMonitorResource_MaximumSet_Gen.json
 */
async function tagRulesListByNewRelicMonitorResourceMaximumSetGen() {
  const subscriptionId = process.env["NEWRELICOBSERVABILITY_SUBSCRIPTION_ID"] || "ddqonpqwjr";
  const resourceGroupName = process.env["NEWRELICOBSERVABILITY_RESOURCE_GROUP"] || "rgopenapi";
  const monitorName = "ipxmlcbonyxtolzejcjshkmlron";
  const credential = new DefaultAzureCredential();
  const client = new NewRelicObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.tagRules.listByNewRelicMonitorResource(
    resourceGroupName,
    monitorName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
