const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or Updates a logger.
 *
 * @summary Creates or Updates a logger.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateAILogger.json
 */
async function apiManagementCreateAiLogger() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const loggerId = "loggerId";
  const parameters = {
    description: "adding a new logger",
    credentials: { instrumentationKey: "11................a1" },
    loggerType: "applicationInsights",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.logger.createOrUpdate(
    resourceGroupName,
    serviceName,
    loggerId,
    parameters
  );
  console.log(result);
}

apiManagementCreateAiLogger().catch(console.error);
