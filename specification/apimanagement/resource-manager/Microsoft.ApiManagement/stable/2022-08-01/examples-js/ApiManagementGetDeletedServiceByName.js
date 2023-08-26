const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get soft-deleted Api Management Service by name.
 *
 * @summary Get soft-deleted Api Management Service by name.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetDeletedServiceByName.json
 */
async function apiManagementGetDeletedServiceByName() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const serviceName = "apimService3";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.deletedServices.getByName(serviceName, location);
  console.log(result);
}
