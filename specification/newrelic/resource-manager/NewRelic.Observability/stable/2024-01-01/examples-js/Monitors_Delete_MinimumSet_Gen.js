const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a NewRelicMonitorResource
 *
 * @summary Delete a NewRelicMonitorResource
 * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/Monitors_Delete_MinimumSet_Gen.json
 */
async function monitorsDeleteMinimumSetGen() {
  const subscriptionId =
    process.env["NEWRELICOBSERVABILITY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["NEWRELICOBSERVABILITY_RESOURCE_GROUP"] || "rgopenapi";
  const userEmail = "ruxvg@xqkmdhrnoo.hlmbpm";
  const monitorName = "ipxmlcbonyxtolzejcjshkmlron";
  const credential = new DefaultAzureCredential();
  const client = new NewRelicObservability(credential, subscriptionId);
  const result = await client.monitors.beginDeleteAndWait(
    resourceGroupName,
    userEmail,
    monitorName,
  );
  console.log(result);
}
