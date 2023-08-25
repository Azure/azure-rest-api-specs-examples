const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Notifies the API Management gateway to create a new connection to the backend after the specified timeout. If no timeout was specified, timeout of 2 minutes is used.
 *
 * @summary Notifies the API Management gateway to create a new connection to the backend after the specified timeout. If no timeout was specified, timeout of 2 minutes is used.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementBackendReconnect.json
 */
async function apiManagementBackendReconnect() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const backendId = "proxybackend";
  const parameters = { after: "PT3S" };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.backend.reconnect(resourceGroupName, serviceName, backendId, options);
  console.log(result);
}
