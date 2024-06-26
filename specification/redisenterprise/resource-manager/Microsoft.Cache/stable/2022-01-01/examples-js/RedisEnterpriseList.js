const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all RedisEnterprise clusters in the specified subscription.
 *
 * @summary Gets all RedisEnterprise clusters in the specified subscription.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseList.json
 */
async function redisEnterpriseList() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.redisEnterprise.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

redisEnterpriseList().catch(console.error);
