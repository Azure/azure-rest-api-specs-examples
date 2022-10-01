const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified logger.
 *
 * @summary Deletes the specified logger.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteLogger.json
 */
async function apiManagementDeleteLogger() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const loggerId = "loggerId";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.logger.delete(resourceGroupName, serviceName, loggerId, ifMatch);
  console.log(result);
}

apiManagementDeleteLogger().catch(console.error);
