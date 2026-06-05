const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get metric status
 *
 * @summary get metric status
 * x-ms-original-file: 2024-04-24/Monitors_GetMetricStatus_MinimumSet_Gen.json
 */
async function monitorsGetMetricStatusMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "1d701e7e-3150-4d33-9279-d4ea03e9110e";
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.monitors.getMetricStatus("rgDynatrace", "fhcjxnxumkdlgpwanewtkdnyuz");
  console.log(result);
}
