const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List NewRelicMonitorResource resources by subscription ID
 *
 * @summary List NewRelicMonitorResource resources by subscription ID
 * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/Monitors_ListBySubscription_MaximumSet_Gen.json
 */
async function monitorsListBySubscriptionMaximumSetGen() {
  const subscriptionId =
    process.env["NEWRELICOBSERVABILITY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new NewRelicObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitors.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
