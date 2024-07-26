const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of access policies associated with this redis cache
 *
 * @summary Gets the list of access policies associated with this redis cache
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCacheAccessPolicyList.json
 */
async function redisCacheAccessPolicyList() {
  const subscriptionId = process.env["REDIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDIS_RESOURCE_GROUP"] || "rg1";
  const cacheName = "cache1";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.accessPolicy.list(resourceGroupName, cacheName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
