const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a MonitorResource
 *
 * @summary update a MonitorResource
 * x-ms-original-file: 2024-04-24/Monitors_Update_MinimumSet_Gen.json
 */
async function monitorsUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.monitors.update("myResourceGroup", "myMonitor", {});
  console.log(result);
}
