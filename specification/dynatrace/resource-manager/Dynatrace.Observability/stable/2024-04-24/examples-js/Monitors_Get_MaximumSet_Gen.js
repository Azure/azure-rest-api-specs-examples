const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a MonitorResource
 *
 * @summary get a MonitorResource
 * x-ms-original-file: 2024-04-24/Monitors_Get_MaximumSet_Gen.json
 */
async function monitorsGetMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.monitors.get("myResourceGroup", "myMonitor");
  console.log(result);
}
