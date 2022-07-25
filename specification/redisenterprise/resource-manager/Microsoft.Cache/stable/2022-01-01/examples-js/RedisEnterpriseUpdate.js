const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing RedisEnterprise cluster
 *
 * @summary Updates an existing RedisEnterprise cluster
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseUpdate.json
 */
async function redisEnterpriseUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cache1";
  const parameters = {
    minimumTlsVersion: "1.2",
    sku: { name: "EnterpriseFlash_F300", capacity: 9 },
    tags: { tag1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.redisEnterprise.beginUpdateAndWait(
    resourceGroupName,
    clusterName,
    parameters
  );
  console.log(result);
}

redisEnterpriseUpdate().catch(console.error);
