const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets list of App Services with Dynatrace PaaS OneAgent enabled
 *
 * @summary Gets list of App Services with Dynatrace PaaS OneAgent enabled
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/Monitors_ListAppServices_MaximumSet_Gen.json
 */
async function monitorsListAppServicesMaximumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitors.listAppServices(resourceGroupName, monitorName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

monitorsListAppServicesMaximumSetGen().catch(console.error);
