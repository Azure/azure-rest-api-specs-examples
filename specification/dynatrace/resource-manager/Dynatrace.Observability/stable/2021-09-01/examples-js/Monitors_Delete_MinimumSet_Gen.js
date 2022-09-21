const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a MonitorResource
 *
 * @summary Delete a MonitorResource
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/Monitors_Delete_MinimumSet_Gen.json
 */
async function monitorsDeleteMinimumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.monitors.beginDeleteAndWait(resourceGroupName, monitorName);
  console.log(result);
}

monitorsDeleteMinimumSetGen().catch(console.error);
