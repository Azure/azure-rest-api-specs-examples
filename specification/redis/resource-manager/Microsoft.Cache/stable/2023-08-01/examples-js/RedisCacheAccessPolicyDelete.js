const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the access policy from a redis cache
 *
 * @summary Deletes the access policy from a redis cache
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheAccessPolicyDelete.json
 */
async function redisCacheAccessPolicyDelete() {
  const subscriptionId = process.env["REDIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDIS_RESOURCE_GROUP"] || "rg1";
  const cacheName = "cache1";
  const accessPolicyName = "accessPolicy1";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.accessPolicy.beginDeleteAndWait(
    resourceGroupName,
    cacheName,
    accessPolicyName
  );
  console.log(result);
}
