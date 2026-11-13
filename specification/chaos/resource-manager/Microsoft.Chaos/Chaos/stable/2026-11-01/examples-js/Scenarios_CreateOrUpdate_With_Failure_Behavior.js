const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a scenario.
 *
 * @summary create or update a scenario.
 * x-ms-original-file: 2026-11-01/Scenarios_CreateOrUpdate_With_Failure_Behavior.json
 */
async function createOrUpdateAScenarioThatContinuesPastSuppressedActionFailures() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.scenarios.createOrUpdate(
    "exampleRG",
    "exampleWorkspace",
    "bestEffortProbeScenario",
    {
      properties: {
        description:
          "Probes dependencies best-effort, then runs a critical shutdown whose failure must stop the run.",
        failureBehavior: "Continue",
        actions: [
          {
            name: "dependencyProbe",
            actionId: "urn:csci:microsoft:agent:networkDisconnect/1.0.0",
            description:
              "Best-effort dependency probe. Inherits the scenario-level Continue behavior, so a failure here does not fail the run.",
            duration: "PT5M",
          },
          {
            name: "criticalShutdown",
            actionId: "urn:csci:microsoft:compute:shutdown/1.0.0",
            description:
              "Critical shutdown. Overrides the scenario default so that a failure stops the run.",
            duration: "PT10M",
            failureBehavior: "Stop",
            runAfter: {
              behavior: "All",
              items: [
                { type: "Action", name: "dependencyProbe", onActionLifecycle: "AnyTerminal" },
              ],
            },
          },
        ],
      },
    },
  );
  console.log(result);
}
