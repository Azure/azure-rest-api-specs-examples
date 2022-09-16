const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Experiment resource.
 *
 * @summary Create or update a Experiment resource.
 * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2022-07-01-preview/examples/CreateOrUpdateAExperiment.json
 */
async function createOrUpdateAExperimentInAResourceGroup() {
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const resourceGroupName = "exampleRG";
  const experimentName = "exampleExperiment";
  const experiment = {
    identity: { type: "SystemAssigned" },
    location: "centraluseuap",
    selectors: [
      {
        type: "List",
        id: "selector1",
        targets: [
          {
            type: "ChaosTarget",
            id: "/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine",
          },
        ],
      },
    ],
    steps: [
      {
        name: "step1",
        branches: [
          {
            name: "branch1",
            actions: [
              {
                name: "urn:csci:provider:providername:Shutdown/1.0",
                type: "continuous",
                duration: "PT10M",
                parameters: [{ key: "abruptShutdown", value: "false" }],
                selectorId: "selector1",
              },
            ],
          },
        ],
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.experiments.createOrUpdate(
    resourceGroupName,
    experimentName,
    experiment
  );
  console.log(result);
}

createOrUpdateAExperimentInAResourceGroup().catch(console.error);
