const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a scenario definition.
 *
 * @summary create or update a scenario definition.
 * x-ms-original-file: 2026-11-01/ScenarioConfigurations_CreateOrUpdate_CreateOnly.json
 */
async function createAScenarioConfigurationOnlyWhenAbsentOrReturn412WhenItAlreadyExists() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.scenarioConfigurations.createOrUpdate(
    "exampleRG",
    "exampleWorkspace",
    "exampleScenario",
    "exampleConfiguration",
    {
      properties: {
        scenarioId:
          "/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/exampleScenario",
      },
    },
    { ifNoneMatch: "*" },
  );
  console.log(result);
}
