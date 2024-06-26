const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the subscriptions that are being monitored by the NewRelic monitor resource
 *
 * @summary Updates the subscriptions that are being monitored by the NewRelic monitor resource
 * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/MonitoredSubscriptions_Delete.json
 */
async function monitorsDeleteMonitoredSubscriptions() {
  const subscriptionId =
    process.env["NEWRELICOBSERVABILITY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName =
    process.env["NEWRELICOBSERVABILITY_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "myMonitor";
  const configurationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new NewRelicObservability(credential, subscriptionId);
  const result = await client.monitoredSubscriptions.beginDeleteAndWait(
    resourceGroupName,
    monitorName,
    configurationName,
  );
  console.log(result);
}
