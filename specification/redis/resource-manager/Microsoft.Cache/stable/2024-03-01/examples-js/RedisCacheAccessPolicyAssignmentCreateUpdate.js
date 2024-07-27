const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Adds the access policy assignment to the specified users
 *
 * @summary Adds the access policy assignment to the specified users
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCacheAccessPolicyAssignmentCreateUpdate.json
 */
async function redisCacheAccessPolicyAssignmentCreateUpdate() {
  const subscriptionId = process.env["REDIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDIS_RESOURCE_GROUP"] || "rg1";
  const cacheName = "cache1";
  const accessPolicyAssignmentName = "accessPolicyAssignmentName1";
  const parameters = {
    accessPolicyName: "accessPolicy1",
    objectId: "6497c918-11ad-41e7-1b0f-7c518a87d0b0",
    objectIdAlias: "TestAADAppRedis",
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.accessPolicyAssignment.beginCreateUpdateAndWait(
    resourceGroupName,
    cacheName,
    accessPolicyAssignmentName,
    parameters,
  );
  console.log(result);
}
