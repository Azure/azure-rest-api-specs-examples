const { MongoClusterManagementClient } = require("@azure/arm-mongocluster");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a mongo cluster. Update overwrites all properties for the resource. To only modify some of the properties, use PATCH.
 *
 * @summary create or update a mongo cluster. Update overwrites all properties for the resource. To only modify some of the properties, use PATCH.
 * x-ms-original-file: 2025-09-01/MongoClusters_CreatePITR_EntraAuth.json
 */
async function createsAMongoClusterResourceFromAPointInTimeRestoreWithMicrosoftEntraIDAuthenticationModeEnabled() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MongoClusterManagementClient(credential, subscriptionId);
  const result = await client.mongoClusters.createOrUpdate("TestResourceGroup", "myMongoCluster", {
    location: "westus2",
    properties: {
      createMode: "PointInTimeRestore",
      authConfig: { allowedModes: ["MicrosoftEntraID"] },
      restoreParameters: {
        pointInTimeUTC: new Date("2023-01-13T20:07:35Z"),
        sourceResourceId:
          "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/myOtherMongoCluster",
      },
    },
  });
  console.log(result);
}
