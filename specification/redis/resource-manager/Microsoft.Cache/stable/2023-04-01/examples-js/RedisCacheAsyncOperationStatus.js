const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to For checking the ongoing status of an operation
 *
 * @summary For checking the ongoing status of an operation
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCacheAsyncOperationStatus.json
 */
async function redisCacheAsyncOperationStatus() {
  const subscriptionId = process.env["REDIS_SUBSCRIPTION_ID"] || "subid";
  const location = "East US";
  const operationId = "c7ba2bf5-5939-4d79-b037-2964ccf097da";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.asyncOperationStatus.get(location, operationId);
  console.log(result);
}
