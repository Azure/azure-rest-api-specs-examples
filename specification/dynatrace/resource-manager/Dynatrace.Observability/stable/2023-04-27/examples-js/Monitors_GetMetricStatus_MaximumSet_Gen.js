const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get metric status
 *
 * @summary Get metric status
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_GetMetricStatus_MaximumSet_Gen.json
 */
async function monitorsGetMetricStatusMaximumSetGen() {
  const subscriptionId = process.env["DYNATRACE_SUBSCRIPTION_ID"] || "nqmcgifgaqlf";
  const resourceGroupName = process.env["DYNATRACE_RESOURCE_GROUP"] || "rgDynatrace";
  const monitorName = "fhcjxnxumkdlgpwanewtkdnyuz";
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.monitors.getMetricStatus(resourceGroupName, monitorName);
  console.log(result);
}
