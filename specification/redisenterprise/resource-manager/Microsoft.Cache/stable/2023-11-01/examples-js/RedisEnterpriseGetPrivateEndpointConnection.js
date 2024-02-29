const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified private endpoint connection associated with the RedisEnterprise cluster.
 *
 * @summary Gets the specified private endpoint connection associated with the RedisEnterprise cluster.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2023-11-01/examples/RedisEnterpriseGetPrivateEndpointConnection.json
 */
async function redisEnterpriseGetPrivateEndpointConnection() {
  const subscriptionId = process.env["REDISENTERPRISE_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDISENTERPRISE_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cache1";
  const privateEndpointConnectionName = "pectest01";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.get(
    resourceGroupName,
    clusterName,
    privateEndpointConnectionName,
  );
  console.log(result);
}
