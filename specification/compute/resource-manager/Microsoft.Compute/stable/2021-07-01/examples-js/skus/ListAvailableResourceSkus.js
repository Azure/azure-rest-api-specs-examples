const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of Microsoft.Compute SKUs available for your Subscription.
 *
 * @summary Gets the list of Microsoft.Compute SKUs available for your Subscription.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/skus/ListAvailableResourceSkus.json
 */
async function listsAllAvailableResourceSkUs() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.resourceSkus.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsAllAvailableResourceSkUs().catch(console.error);
