const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes the specified operation in the API.
 *
 * @summary Deletes the specified operation in the API.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementDeleteApiOperation.json
 */
async function apiManagementDeleteApiOperation() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "57d2ef278aa04f0888cba3f3";
  const operationId = "57d2ef278aa04f0ad01d6cdc";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiOperation.delete(
    resourceGroupName,
    serviceName,
    apiId,
    operationId,
    ifMatch,
  );
  console.log(result);
}
