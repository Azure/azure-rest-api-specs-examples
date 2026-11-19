const { MongoClusterManagementClient } = require("@azure/arm-mongocluster");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing DocumentDB cluster. The request body can contain one to many of the properties present in the normal DocumentDB cluster definition.
 *
 * @summary updates an existing DocumentDB cluster. The request body can contain one to many of the properties present in the normal DocumentDB cluster definition.
 * x-ms-original-file: 2026-09-01/MongoClusters_PatchTags.json
 */
async function updatesStringValuedResourceTagsWithoutChangingTheExistingPremiumSSDStorageType() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MongoClusterManagementClient(credential, subscriptionId);
  const result = await client.mongoClusters.update("TestResourceGroup", "myMongoCluster", {
    tags: { environment: "production", costCenter: "1234" },
  });
  console.log(result);
}
