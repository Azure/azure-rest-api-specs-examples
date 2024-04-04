const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List FleetUpdateStrategy resources by Fleet
 *
 * @summary List FleetUpdateStrategy resources by Fleet
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-02-02-preview/examples/UpdateStrategies_ListByFleet.json
 */
async function listTheFleetUpdateStrategyResourcesByFleet() {
  const subscriptionId =
    process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const fleetName = "fleet1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fleetUpdateStrategies.listByFleet(resourceGroupName, fleetName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
