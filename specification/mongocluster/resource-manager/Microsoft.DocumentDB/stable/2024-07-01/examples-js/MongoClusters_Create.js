const { MongoClusterManagementClient } = require("@azure/arm-mongocluster");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a mongo cluster. Update overwrites all properties for the resource. To only modify some of the properties, use PATCH.
 *
 * @summary create or update a mongo cluster. Update overwrites all properties for the resource. To only modify some of the properties, use PATCH.
 * x-ms-original-file: 2024-07-01/MongoClusters_Create.json
 */
async function createsANewMongoClusterResource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MongoClusterManagementClient(credential, subscriptionId);
  const result = await client.mongoClusters.createOrUpdate("TestResourceGroup", "myMongoCluster", {
    location: "westus2",
    properties: {
      administrator: { userName: "mongoAdmin", password: "password" },
      serverVersion: "5.0",
      storage: { sizeGb: 128 },
      compute: { tier: "M30" },
      sharding: { shardCount: 1 },
      highAvailability: { targetMode: "SameZone" },
    },
  });
  console.log(result);
}
