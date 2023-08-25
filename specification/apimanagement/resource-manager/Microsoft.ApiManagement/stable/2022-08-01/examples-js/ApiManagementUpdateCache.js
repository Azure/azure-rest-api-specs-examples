const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the details of the cache specified by its identifier.
 *
 * @summary Updates the details of the cache specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateCache.json
 */
async function apiManagementUpdateCache() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const cacheId = "c1";
  const ifMatch = "*";
  const parameters = { useFromLocation: "westindia" };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.cache.update(
    resourceGroupName,
    serviceName,
    cacheId,
    ifMatch,
    parameters
  );
  console.log(result);
}
