const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a UpdateRun
 *
 * @summary delete a UpdateRun
 * x-ms-original-file: 2025-04-01-preview/UpdateRuns_Delete_MaximumSet_Gen.json
 */
async function deleteAnUpdateRunResourceGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  await client.updateRuns.delete("rgfleets", "fleet1", "fleet1", {
    ifMatch: "xnbwucfeufeagpa",
  });
}
