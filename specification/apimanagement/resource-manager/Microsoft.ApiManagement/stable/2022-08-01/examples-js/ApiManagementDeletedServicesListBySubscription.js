const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all soft-deleted services available for undelete for the given subscription.
 *
 * @summary Lists all soft-deleted services available for undelete for the given subscription.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeletedServicesListBySubscription.json
 */
async function apiManagementDeletedServicesListBySubscription() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.deletedServices.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
