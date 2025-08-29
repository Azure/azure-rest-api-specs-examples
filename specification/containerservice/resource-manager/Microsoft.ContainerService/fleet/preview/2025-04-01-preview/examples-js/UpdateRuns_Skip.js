const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

async function skipsOneOrMoreMemberOrGroupOrStageOrAfterStageWaitSOfAnUpdateRun() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.updateRuns.skip("rg1", "fleet1", "run1", {
    targets: [
      { type: "Member", name: "member-one" },
      { type: "AfterStageWait", name: "stage1" },
    ],
  });
  console.log(result);
}
