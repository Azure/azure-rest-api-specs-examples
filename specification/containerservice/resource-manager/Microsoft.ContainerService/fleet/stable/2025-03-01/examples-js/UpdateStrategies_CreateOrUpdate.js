const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a FleetUpdateStrategy
 *
 * @summary create a FleetUpdateStrategy
 * x-ms-original-file: 2025-03-01/UpdateStrategies_CreateOrUpdate.json
 */
async function createAFleetUpdateStrategy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.fleetUpdateStrategies.createOrUpdate("rg1", "fleet1", "strartegy1", {
    properties: {
      strategy: {
        stages: [
          {
            name: "stage1",
            groups: [{ name: "group-a" }],
            afterStageWaitInSeconds: 3600,
          },
        ],
      },
    },
  });
  console.log(result);
}
