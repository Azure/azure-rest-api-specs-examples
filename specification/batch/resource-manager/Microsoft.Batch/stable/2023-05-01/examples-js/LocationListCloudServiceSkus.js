const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of Batch supported Cloud Service VM sizes available at the given location.
 *
 * @summary Gets the list of Batch supported Cloud Service VM sizes available at the given location.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/LocationListCloudServiceSkus.json
 */
async function locationListCloudServiceSkus() {
  const subscriptionId = process.env["BATCH_SUBSCRIPTION_ID"] || "subid";
  const locationName = "japaneast";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.location.listSupportedCloudServiceSkus(locationName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
