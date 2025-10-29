const { MongoClusterManagementClient } = require("@azure/arm-mongocluster");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a mongo cluster. Update overwrites all properties for the resource. To only modify some of the properties, use PATCH.
 *
 * @summary create or update a mongo cluster. Update overwrites all properties for the resource. To only modify some of the properties, use PATCH.
 * x-ms-original-file: 2025-09-01/MongoClusters_Create_SSDv2.json
 */
async function createsANewMongoClusterResourceWithPremiumSSDv2Storage() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MongoClusterManagementClient(credential, subscriptionId);
  const result = await client.mongoClusters.createOrUpdate("TestResourceGroup", "myMongoCluster", {
    location: "westus2",
    properties: {
      administrator: { userName: "mongoAdmin", password: "password" },
      authConfig: { allowedModes: ["NativeAuth"] },
      serverVersion: "5.0",
      storage: { sizeGb: 32, type: "PremiumSSDv2" },
      compute: { tier: "M30" },
      sharding: { shardCount: 1 },
      highAvailability: { targetMode: "ZoneRedundantPreferred" },
    },
  });
  console.log(result);
}
