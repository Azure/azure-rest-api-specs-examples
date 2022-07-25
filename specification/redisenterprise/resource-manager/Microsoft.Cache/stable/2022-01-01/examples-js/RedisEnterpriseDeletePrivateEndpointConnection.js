const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified private endpoint connection associated with the RedisEnterprise cluster.
 *
 * @summary Deletes the specified private endpoint connection associated with the RedisEnterprise cluster.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDeletePrivateEndpointConnection.json
 */
async function redisEnterpriseDeletePrivateEndpointConnection() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cache1";
  const privateEndpointConnectionName = "pectest01";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.delete(
    resourceGroupName,
    clusterName,
    privateEndpointConnectionName
  );
  console.log(result);
}

redisEnterpriseDeletePrivateEndpointConnection().catch(console.error);
