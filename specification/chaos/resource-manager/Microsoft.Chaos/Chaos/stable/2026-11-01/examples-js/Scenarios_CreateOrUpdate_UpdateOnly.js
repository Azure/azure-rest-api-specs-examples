const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a scenario.
 *
 * @summary create or update a scenario.
 * x-ms-original-file: 2026-11-01/Scenarios_CreateOrUpdate_UpdateOnly.json
 */
async function updateAScenarioOnlyWhenPresentOrReturn412WhenItDoesNotExist() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.scenarios.createOrUpdate(
    "exampleRG",
    "exampleWorkspace",
    "exampleScenario",
    {
      properties: {
        parameters: [],
        actions: [
          {
            name: "restart",
            actionId: "microsoft-compute-shutdown/2.0",
            duration: "PT5M",
            selectors: [
              { name: "logicalZone", value: "1" },
              { name: "location", value: "eastus" },
            ],
          },
        ],
      },
    },
    { ifMatch: "*" },
  );
  console.log(result);
}
