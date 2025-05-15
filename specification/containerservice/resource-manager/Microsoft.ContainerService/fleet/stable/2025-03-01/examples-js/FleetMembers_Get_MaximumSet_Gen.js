const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a FleetMember
 *
 * @summary get a FleetMember
 * x-ms-original-file: 2025-03-01/FleetMembers_Get_MaximumSet_Gen.json
 */
async function getsAFleetMemberResourceGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.fleetMembers.get("rgfleets", "fleet1", "fleet1");
  console.log(result);
}
