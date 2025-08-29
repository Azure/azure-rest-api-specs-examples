const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

async function createAnUpdateRun() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.updateRuns.createOrUpdate("rg1", "fleet1", "run1", {
    properties: {
      updateStrategyId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/updateStrategies/strategy1",
      strategy: {
        stages: [
          {
            name: "stage1",
            groups: [
              {
                name: "group-a",
                beforeGates: [{ displayName: "gate before group-a", type: "Approval" }],
                afterGates: [{ displayName: "gate after group-a", type: "Approval" }],
              },
            ],
            beforeGates: [{ displayName: "gate before stage1", type: "Approval" }],
            afterGates: [{ displayName: "gate after stage1", type: "Approval" }],
            afterStageWaitInSeconds: 3600,
          },
        ],
      },
      managedClusterUpdate: {
        upgrade: { type: "Full", kubernetesVersion: "1.26.1" },
        nodeImageSelection: { type: "Latest" },
      },
    },
  });
  console.log(result);
}
