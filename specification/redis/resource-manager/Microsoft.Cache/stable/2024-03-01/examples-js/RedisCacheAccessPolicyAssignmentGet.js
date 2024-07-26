const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of assignments for an access policy of a redis cache
 *
 * @summary Gets the list of assignments for an access policy of a redis cache
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCacheAccessPolicyAssignmentGet.json
 */
async function redisCacheAccessPolicyAssignmentGet() {
  const subscriptionId = process.env["REDIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDIS_RESOURCE_GROUP"] || "rg1";
  const cacheName = "cache1";
  const accessPolicyAssignmentName = "accessPolicyAssignmentName1";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.accessPolicyAssignment.get(
    resourceGroupName,
    cacheName,
    accessPolicyAssignmentName,
  );
  console.log(result);
}
