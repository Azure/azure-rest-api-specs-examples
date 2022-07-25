const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all Virtual Instances for SAP in the subscription.
 *
 * @summary Gets all Virtual Instances for SAP in the subscription.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_ListBySubscription.json
 */
async function sapVirtualInstancesListBySubscription() {
  const subscriptionId = "6d875e77-e412-4d7d-9af4-8895278b4443";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sAPVirtualInstances.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

sapVirtualInstancesListBySubscription().catch(console.error);
