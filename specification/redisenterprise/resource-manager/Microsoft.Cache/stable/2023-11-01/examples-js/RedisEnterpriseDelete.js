const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a RedisEnterprise cache cluster.
 *
 * @summary Deletes a RedisEnterprise cache cluster.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2023-11-01/examples/RedisEnterpriseDelete.json
 */
async function redisEnterpriseDelete() {
  const subscriptionId = process.env["REDISENTERPRISE_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDISENTERPRISE_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cache1";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.redisEnterprise.beginDeleteAndWait(resourceGroupName, clusterName);
  console.log(result);
}
