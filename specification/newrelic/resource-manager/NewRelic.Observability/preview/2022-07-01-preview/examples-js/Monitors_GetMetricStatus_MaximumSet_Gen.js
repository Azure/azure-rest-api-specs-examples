const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get metric status
 *
 * @summary Get metric status
 * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/preview/2022-07-01-preview/examples/Monitors_GetMetricStatus_MaximumSet_Gen.json
 */
async function monitorsGetMetricStatusMaximumSetGen() {
  const subscriptionId = process.env["NEWRELICOBSERVABILITY_SUBSCRIPTION_ID"] || "nqmcgifgaqlf";
  const resourceGroupName = process.env["NEWRELICOBSERVABILITY_RESOURCE_GROUP"] || "rgNewRelic";
  const monitorName = "fhcjxnxumkdlgpwanewtkdnyuz";
  const request = {
    azureResourceIds: ["enfghpfw"],
    userEmail: "ruxvg@xqkmdhrnoo.hlmbpm",
  };
  const credential = new DefaultAzureCredential();
  const client = new NewRelicObservability(credential, subscriptionId);
  const result = await client.monitors.getMetricStatus(resourceGroupName, monitorName, request);
  console.log(result);
}
