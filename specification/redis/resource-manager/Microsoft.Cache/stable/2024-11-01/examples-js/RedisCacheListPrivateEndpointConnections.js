const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List all the private endpoint connections associated with the redis cache.
 *
 * @summary List all the private endpoint connections associated with the redis cache.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/RedisCacheListPrivateEndpointConnections.json
 */
async function redisCacheListPrivateEndpointConnection() {
  const subscriptionId = process.env["REDIS_SUBSCRIPTION_ID"] || "{subscriptionId}";
  const resourceGroupName = process.env["REDIS_RESOURCE_GROUP"] || "rgtest01";
  const cacheName = "cachetest01";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.list(resourceGroupName, cacheName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
