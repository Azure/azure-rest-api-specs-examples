const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a RedisEnterprise cache cluster.
 *
 * @summary Deletes a RedisEnterprise cache cluster.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDelete.json
 */
async function redisEnterpriseDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cache1";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.redisEnterprise.beginDeleteAndWait(resourceGroupName, clusterName);
  console.log(result);
}

redisEnterpriseDelete().catch(console.error);
