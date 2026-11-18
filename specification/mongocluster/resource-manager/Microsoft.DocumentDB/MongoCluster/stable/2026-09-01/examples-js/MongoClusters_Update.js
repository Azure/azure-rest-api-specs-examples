const { MongoClusterManagementClient } = require("@azure/arm-mongocluster");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing DocumentDB cluster. The request body can contain one to many of the properties present in the normal DocumentDB cluster definition.
 *
 * @summary updates an existing DocumentDB cluster. The request body can contain one to many of the properties present in the normal DocumentDB cluster definition.
 * x-ms-original-file: 2026-09-01/MongoClusters_Update.json
 */
async function updatesADocumentDBClusterResource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MongoClusterManagementClient(credential, subscriptionId);
  const result = await client.mongoClusters.update("TestResourceGroup", "myMongoCluster", {
    properties: {
      administrator: { userName: "mongoAdmin" },
      authConfig: { allowedModes: ["NativeAuth"] },
      serverVersion: "5.0",
      storage: { sizeGb: 256, type: "PremiumSSDv2" },
      compute: { tier: "M50" },
      sharding: { shardCount: 4 },
      highAvailability: { targetMode: "SameZone" },
      previewFeatures: [],
      publicNetworkAccess: "Enabled",
      dataApi: { mode: "Disabled" },
    },
  });
  console.log(result);
}
