const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a MonitorResource
 *
 * @summary create a MonitorResource
 * x-ms-original-file: 2024-04-24/Monitors_CreateOrUpdate_MinimumSet_Gen.json
 */
async function monitorsCreateOrUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.monitors.createOrUpdate("myResourceGroup", "myMonitor", {
    location: "West US 2",
  });
  console.log(result);
}
