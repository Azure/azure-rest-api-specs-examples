const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the list of Microsoft.Compute SKUs available for your Subscription.
 *
 * @summary gets the list of Microsoft.Compute SKUs available for your Subscription.
 * x-ms-original-file: 2021-07-01/skus/ListAvailableResourceSkusWithExtendedLocations.json
 */
async function listsAllAvailableResourceSKUsWithExtendedLocationInformation() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.resourceSkus.list({ includeExtendedLocations: "true" })) {
    resArray.push(item);
  }

  console.log(resArray);
}
