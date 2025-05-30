const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or Updates a logger.
 *
 * @summary Creates or Updates a logger.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateWorkspaceAILogger.json
 */
async function apiManagementCreateWorkspaceAiLogger() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const workspaceId = "wks1";
  const loggerId = "loggerId";
  const parameters = {
    description: "adding a new logger",
    credentials: { instrumentationKey: "11................a1" },
    loggerType: "applicationInsights",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.workspaceLogger.createOrUpdate(
    resourceGroupName,
    serviceName,
    workspaceId,
    loggerId,
    parameters,
  );
  console.log(result);
}
