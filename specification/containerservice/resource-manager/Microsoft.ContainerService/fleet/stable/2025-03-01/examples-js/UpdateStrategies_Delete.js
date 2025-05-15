const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a FleetUpdateStrategy
 *
 * @summary delete a FleetUpdateStrategy
 * x-ms-original-file: 2025-03-01/UpdateStrategies_Delete.json
 */
async function deleteAFleetUpdateStrategyResource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  await client.fleetUpdateStrategies.delete("rg1", "fleet1", "strategy1");
}
