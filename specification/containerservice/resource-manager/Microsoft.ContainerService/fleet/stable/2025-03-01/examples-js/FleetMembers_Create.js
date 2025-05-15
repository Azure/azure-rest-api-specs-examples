const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a FleetMember
 *
 * @summary create a FleetMember
 * x-ms-original-file: 2025-03-01/FleetMembers_Create.json
 */
async function createsAFleetMemberResourceWithALongRunningOperation() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.fleetMembers.create("rg1", "fleet1", "member-1", {
    properties: {
      clusterResourceId:
        "/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster-1",
    },
  });
  console.log(result);
}
