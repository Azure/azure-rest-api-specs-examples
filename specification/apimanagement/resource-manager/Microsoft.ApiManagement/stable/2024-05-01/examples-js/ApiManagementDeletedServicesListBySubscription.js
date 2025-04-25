const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all soft-deleted services available for undelete for the given subscription.
 *
 * @summary Lists all soft-deleted services available for undelete for the given subscription.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementDeletedServicesListBySubscription.json
 */
async function apiManagementDeletedServicesListBySubscription() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.deletedServices.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
