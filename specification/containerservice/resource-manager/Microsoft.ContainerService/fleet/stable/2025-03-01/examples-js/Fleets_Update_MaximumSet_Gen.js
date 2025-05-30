const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a Fleet
 *
 * @summary update a Fleet
 * x-ms-original-file: 2025-03-01/Fleets_Update_MaximumSet_Gen.json
 */
async function updateAFleetGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.fleets.updateAsync(
    "rgfleets",
    "fleet1",
    {
      tags: {},
      identity: { type: "None", userAssignedIdentities: { key126: {} } },
    },
    { ifMatch: "lgoeir" },
  );
  console.log(result);
}
