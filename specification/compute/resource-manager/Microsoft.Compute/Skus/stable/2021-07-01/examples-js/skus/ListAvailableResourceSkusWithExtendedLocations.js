const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of Microsoft.Compute SKUs available for your Subscription.
 *
 * @summary Gets the list of Microsoft.Compute SKUs available for your Subscription.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/Skus/stable/2021-07-01/examples/skus/ListAvailableResourceSkusWithExtendedLocations.json
 */
async function listsAllAvailableResourceSkUsWithExtendedLocationInformation() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const includeExtendedLocations = "true";
  const options = { includeExtendedLocations };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.resourceSkus.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
