const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update a managed Cassandra cluster. When updating, you must specify all writable properties. To update only some properties, use PATCH.
 *
 * @summary Create or update a managed Cassandra cluster. When updating, you must specify all writable properties. To update only some properties, use PATCH.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBManagedCassandraClusterCreate.json
 */
async function cosmosDbManagedCassandraClusterCreate() {
  const subscriptionId =
    process.env["COSMOSDB_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "cassandra-prod-rg";
  const clusterName = "cassandra-prod";
  const body = {
    location: "West US",
    properties: {
      authenticationMethod: "Cassandra",
      cassandraVersion: "3.11",
      clientCertificates: [
        {
          pem: "-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----",
        },
      ],
      clusterNameOverride: "ClusterNameIllegalForAzureResource",
      delegatedManagementSubnetId:
        "/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management",
      externalGossipCertificates: [
        {
          pem: "-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----",
        },
      ],
      externalSeedNodes: [
        { ipAddress: "10.52.221.2" },
        { ipAddress: "10.52.221.3" },
        { ipAddress: "10.52.221.4" },
      ],
      hoursBetweenBackups: 24,
      initialCassandraAdminPassword: "mypassword",
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.cassandraClusters.beginCreateUpdateAndWait(
    resourceGroupName,
    clusterName,
    body,
  );
  console.log(result);
}
