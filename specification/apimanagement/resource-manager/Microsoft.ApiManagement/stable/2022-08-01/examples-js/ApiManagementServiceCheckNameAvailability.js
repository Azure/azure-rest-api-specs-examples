const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks availability and correctness of a name for an API Management service.
 *
 * @summary Checks availability and correctness of a name for an API Management service.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceCheckNameAvailability.json
 */
async function apiManagementServiceCheckNameAvailability() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const parameters = {
    name: "apimService1",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.checkNameAvailability(parameters);
  console.log(result);
}
