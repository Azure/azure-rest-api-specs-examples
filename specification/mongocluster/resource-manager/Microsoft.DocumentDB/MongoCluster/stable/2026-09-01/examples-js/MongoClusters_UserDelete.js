const { MongoClusterManagementClient } = require("@azure/arm-mongocluster");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes a DocumentDB cluster user.
 *
 * @summary deletes a DocumentDB cluster user.
 * x-ms-original-file: 2026-09-01/MongoClusters_UserDelete.json
 */
async function deletesAUserOnADocumentDBClusterResource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MongoClusterManagementClient(credential, subscriptionId);
  await client.users.delete("TestGroup", "myMongoCluster", "uuuuuuuu-uuuu-uuuu-uuuu-uuuuuuuuuuuu");
}
