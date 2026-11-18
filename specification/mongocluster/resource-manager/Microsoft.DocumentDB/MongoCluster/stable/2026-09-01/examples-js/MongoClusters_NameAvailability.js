const { MongoClusterManagementClient } = require("@azure/arm-mongocluster");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to check if DocumentDB cluster name is available for use.
 *
 * @summary check if DocumentDB cluster name is available for use.
 * x-ms-original-file: 2026-09-01/MongoClusters_NameAvailability.json
 */
async function checksAndConfirmsTheDocumentDBClusterNameIsAvailableForUse() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MongoClusterManagementClient(credential, subscriptionId);
  const result = await client.mongoClusters.checkNameAvailability("westus2", {
    name: "newmongocluster",
    type: "Microsoft.DocumentDB/mongoClusters",
  });
  console.log(result);
}
