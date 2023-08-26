const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an External Cache to be used in Api Management instance.
 *
 * @summary Creates or updates an External Cache to be used in Api Management instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateCache.json
 */
async function apiManagementCreateCache() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const cacheId = "c1";
  const parameters = {
    description: "Redis cache instances in West India",
    connectionString: "apim.redis.cache.windows.net:6380,password=xc,ssl=True,abortConnect=False",
    resourceId:
      "https://management.azure.com/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redis/apimservice1",
    useFromLocation: "default",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.cache.createOrUpdate(
    resourceGroupName,
    serviceName,
    cacheId,
    parameters
  );
  console.log(result);
}
