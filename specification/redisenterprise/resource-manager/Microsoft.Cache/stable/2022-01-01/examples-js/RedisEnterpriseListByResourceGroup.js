const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all RedisEnterprise clusters in a resource group.
 *
 * @summary Lists all RedisEnterprise clusters in a resource group.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseListByResourceGroup.json
 */
async function redisEnterpriseListByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.redisEnterprise.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

redisEnterpriseListByResourceGroup().catch(console.error);
