const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a FleetUpdateStrategy
 *
 * @summary Delete a FleetUpdateStrategy
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2024-04-01/examples/UpdateStrategies_Delete.json
 */
async function deleteAFleetUpdateStrategyResource() {
  const subscriptionId = process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "subid1";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const fleetName = "fleet1";
  const updateStrategyName = "strategy1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.fleetUpdateStrategies.beginDeleteAndWait(
    resourceGroupName,
    fleetName,
    updateStrategyName,
  );
  console.log(result);
}
