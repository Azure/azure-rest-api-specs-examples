const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to skips one or a combination of member/group/stage/afterStageWait(s) of an update run.
 *
 * @summary skips one or a combination of member/group/stage/afterStageWait(s) of an update run.
 * x-ms-original-file: 2025-03-01/UpdateRuns_Skip_MaximumSet_Gen.json
 */
async function skipsOneOrMoreMemberOrGroupOrStageOrAfterStageWaitSOfAnUpdateRunGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.updateRuns.skip(
    "rgfleets",
    "fleet1",
    "fleet1",
    {
      targets: [
        { type: "Member", name: "member-one" },
        { type: "AfterStageWait", name: "stage1" },
      ],
    },
    { ifMatch: "rncfubdzrhcihvpqflbsjvoau" },
  );
  console.log(result);
}
