const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all MonitorResource by subscriptionId
 *
 * @summary List all MonitorResource by subscriptionId
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/Monitors_ListBySubscriptionId_MinimumSet_Gen.json
 */
async function monitorsListBySubscriptionIdMinimumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitors.listBySubscriptionId()) {
    resArray.push(item);
  }
  console.log(resArray);
}

monitorsListBySubscriptionIdMinimumSetGen().catch(console.error);
