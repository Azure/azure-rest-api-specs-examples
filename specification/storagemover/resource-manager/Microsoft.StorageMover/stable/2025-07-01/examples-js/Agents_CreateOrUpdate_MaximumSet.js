const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an Agent resource, which references a hybrid compute machine that can run jobs.
 *
 * @summary creates or updates an Agent resource, which references a hybrid compute machine that can run jobs.
 * x-ms-original-file: 2025-07-01/Agents_CreateOrUpdate_MaximumSet.json
 */
async function agentsCreateOrUpdateMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.agents.createOrUpdate(
    "examples-rg",
    "examples-storageMoverName",
    "examples-agentName",
    {
      properties: {
        description: "Example Agent Description",
        arcResourceId:
          "/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName",
        arcVmUuid: "3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9",
        uploadLimitSchedule: {
          weeklyRecurrences: [
            {
              days: ["Monday"],
              endTime: { hour: 18, minute: 30 },
              limitInMbps: 2000,
              startTime: { hour: 9, minute: 0 },
            },
          ],
        },
      },
    },
  );
  console.log(result);
}
