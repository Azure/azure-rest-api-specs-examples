const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a scenario.
 *
 * @summary create or update a scenario.
 * x-ms-original-file: 2026-11-01/Scenarios_CreateOrUpdate_From_Template.json
 */
async function createAScenarioFromABuiltInScenarioTemplateVersion() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.scenarios.createOrUpdate(
    "exampleRG",
    "exampleWorkspace",
    "computeZoneDownScenario",
    {
      properties: {
        createdFrom: "/providers/Microsoft.Chaos/scenarioTemplates/ComputeZoneDown/versions/1.2",
        description:
          "Simulates a datacenter/availability zone failure by shutting down VMSS instances and standalone VMs in the target zone. Zone failures account for ~25% of all cloud incidents — this scenario validates your application's cross-zone resilience.",
      },
    },
  );
  console.log(result);
}
