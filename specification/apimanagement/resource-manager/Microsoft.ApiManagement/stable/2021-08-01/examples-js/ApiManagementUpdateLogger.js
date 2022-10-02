const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing logger.
 *
 * @summary Updates an existing logger.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateLogger.json
 */
async function apiManagementUpdateLogger() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const loggerId = "eh1";
  const ifMatch = "*";
  const parameters = {
    description: "updating description",
    loggerType: "azureEventHub",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.logger.update(
    resourceGroupName,
    serviceName,
    loggerId,
    ifMatch,
    parameters
  );
  console.log(result);
}

apiManagementUpdateLogger().catch(console.error);
