const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or Updates a logger.
 *
 * @summary Creates or Updates a logger.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateEHLogger.json
 */
async function apiManagementCreateEhLogger() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const loggerId = "eh1";
  const parameters = {
    description: "adding a new logger",
    credentials: {
      name: "hydraeventhub",
      connectionString:
        "Endpoint=sb://hydraeventhub-ns.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=********=",
    },
    loggerType: "azureEventHub",
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
