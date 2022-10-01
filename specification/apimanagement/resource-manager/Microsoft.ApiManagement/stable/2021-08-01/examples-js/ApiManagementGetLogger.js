const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of the logger specified by its identifier.
 *
 * @summary Gets the details of the logger specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetLogger.json
 */
async function apiManagementGetLogger() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const loggerId = "templateLogger";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.logger.get(resourceGroupName, serviceName, loggerId);
  console.log(result);
}

apiManagementGetLogger().catch(console.error);
