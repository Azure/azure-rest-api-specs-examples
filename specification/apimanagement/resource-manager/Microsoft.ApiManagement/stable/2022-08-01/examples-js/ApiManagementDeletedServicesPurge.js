const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Purges Api Management Service (deletes it with no option to undelete).
 *
 * @summary Purges Api Management Service (deletes it with no option to undelete).
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeletedServicesPurge.json
 */
async function apiManagementDeletedServicesPurge() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const serviceName = "apimService3";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.deletedServices.beginPurgeAndWait(serviceName, location);
  console.log(result);
}
