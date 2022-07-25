const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the private endpoint connections associated with the RedisEnterprise cluster.
 *
 * @summary Lists all the private endpoint connections associated with the RedisEnterprise cluster.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseListPrivateEndpointConnections.json
 */
async function redisEnterpriseListPrivateEndpointConnections() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cache1";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.list(resourceGroupName, clusterName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

redisEnterpriseListPrivateEndpointConnections().catch(console.error);
