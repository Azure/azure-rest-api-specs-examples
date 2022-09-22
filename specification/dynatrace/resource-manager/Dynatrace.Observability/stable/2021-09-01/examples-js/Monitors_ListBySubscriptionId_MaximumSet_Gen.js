const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all MonitorResource by subscriptionId
 *
 * @summary List all MonitorResource by subscriptionId
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/Monitors_ListBySubscriptionId_MaximumSet_Gen.json
 */
async function monitorsListBySubscriptionIdMaximumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitors.listBySubscriptionId()) {
    resArray.push(item);
  }
  console.log(resArray);
}

monitorsListBySubscriptionIdMaximumSetGen().catch(console.error);
