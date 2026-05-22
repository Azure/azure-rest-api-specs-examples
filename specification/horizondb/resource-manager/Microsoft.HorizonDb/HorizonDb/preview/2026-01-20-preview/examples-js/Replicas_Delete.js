const { HorizonDbClient } = require("@azure/arm-horizondb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes a HorizonDb replica.
 *
 * @summary deletes a HorizonDb replica.
 * x-ms-original-file: 2026-01-20-preview/Replicas_Delete.json
 */
async function deleteAHorizonDbReplica() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new HorizonDbClient(credential, subscriptionId);
  await client.horizonDbReplicas.delete(
    "exampleresourcegroup",
    "examplecluster",
    "examplepool",
    "examplereplica",
  );
}
