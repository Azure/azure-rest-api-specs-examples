const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a database
 *
 * @summary Creates a database
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDatabasesCreateWithGeoReplication.json
 */
async function redisEnterpriseDatabasesCreateWithActiveGeoReplication() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const clusterName = "cache1";
  const databaseName = "default";
  const parameters = {
    clientProtocol: "Encrypted",
    clusteringPolicy: "EnterpriseCluster",
    evictionPolicy: "NoEviction",
    geoReplication: {
      groupNickname: "groupName",
      linkedDatabases: [
        {
          id: "/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default",
        },
        {
          id: "/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default",
        },
      ],
    },
    port: 10000,
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.databases.beginCreateAndWait(
    resourceGroupName,
    clusterName,
    databaseName,
    parameters
  );
  console.log(result);
}

redisEnterpriseDatabasesCreateWithActiveGeoReplication().catch(console.error);
