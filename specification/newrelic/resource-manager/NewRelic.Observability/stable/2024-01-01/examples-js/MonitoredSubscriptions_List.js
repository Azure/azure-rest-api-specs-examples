const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the subscriptions currently being monitored by the NewRelic monitor resource.
 *
 * @summary List the subscriptions currently being monitored by the NewRelic monitor resource.
 * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/MonitoredSubscriptions_List.json
 */
async function monitorsGetMonitoredSubscriptions() {
  const subscriptionId =
    process.env["NEWRELICOBSERVABILITY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName =
    process.env["NEWRELICOBSERVABILITY_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new NewRelicObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitoredSubscriptions.list(resourceGroupName, monitorName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
