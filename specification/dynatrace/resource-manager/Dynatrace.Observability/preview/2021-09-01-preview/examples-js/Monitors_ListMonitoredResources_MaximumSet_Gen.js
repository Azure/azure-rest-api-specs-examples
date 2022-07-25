const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the resources currently being monitored by the Dynatrace monitor resource.
 *
 * @summary List the resources currently being monitored by the Dynatrace monitor resource.
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/Monitors_ListMonitoredResources_MaximumSet_Gen.json
 */
async function monitorsListMonitoredResourcesMaximumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitors.listMonitoredResources(resourceGroupName, monitorName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

monitorsListMonitoredResourcesMaximumSetGen().catch(console.error);
