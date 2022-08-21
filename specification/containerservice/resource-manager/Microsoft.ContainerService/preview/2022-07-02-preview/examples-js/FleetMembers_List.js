const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the members of a fleet.
 *
 * @summary Lists the members of a fleet.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-07-02-preview/examples/FleetMembers_List.json
 */
async function listsTheMembersOfAFleet() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const fleetName = "fleet-1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fleetMembers.listByFleet(resourceGroupName, fleetName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsTheMembersOfAFleet().catch(console.error);
