const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Returns the current status of an async operation.
 *
 * @summary Returns the current status of an async operation.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetOperationStatus.json
 */
async function getOperationStatus() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const location = "testLocation";
  const operationId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.operationStatus.get(location, operationId);
  console.log(result);
}
