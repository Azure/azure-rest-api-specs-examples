const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of SAP monitors in the specified subscription. The operations returns various properties of each SAP monitor.
 *
 * @summary Gets a list of SAP monitors in the specified subscription. The operations returns various properties of each SAP monitor.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/workloadmonitor/monitors_List.json
 */
async function listAllSapMonitorsInASubscription() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitors.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
