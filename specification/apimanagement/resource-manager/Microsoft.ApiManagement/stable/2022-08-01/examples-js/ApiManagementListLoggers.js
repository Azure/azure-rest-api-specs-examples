const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists a collection of loggers in the specified service instance.
 *
 * @summary Lists a collection of loggers in the specified service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListLoggers.json
 */
async function apiManagementListLoggers() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.logger.listByService(resourceGroupName, serviceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
