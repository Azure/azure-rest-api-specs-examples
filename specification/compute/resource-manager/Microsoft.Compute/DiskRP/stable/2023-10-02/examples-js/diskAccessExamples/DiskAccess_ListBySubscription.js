const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the disk access resources under a subscription.
 *
 * @summary Lists all the disk access resources under a subscription.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/diskAccessExamples/DiskAccess_ListBySubscription.json
 */
async function listAllDiskAccessResourcesInASubscription() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.diskAccesses.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
