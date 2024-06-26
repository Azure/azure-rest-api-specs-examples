const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all Virtual Instances for SAP solutions resources in a Subscription.
 *
 * @summary Gets all Virtual Instances for SAP solutions resources in a Subscription.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_ListBySubscription.json
 */
async function sapVirtualInstancesListBySubscription() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "6d875e77-e412-4d7d-9af4-8895278b4443";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sAPVirtualInstances.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
