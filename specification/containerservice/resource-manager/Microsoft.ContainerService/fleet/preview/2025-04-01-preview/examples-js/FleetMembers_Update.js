const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

async function updatesAFleetMemberResourceSynchronously() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.fleetMembers.updateAsync("rg1", "fleet1", "member-1", {
    properties: { group: "staging" },
  });
  console.log(result);
}
