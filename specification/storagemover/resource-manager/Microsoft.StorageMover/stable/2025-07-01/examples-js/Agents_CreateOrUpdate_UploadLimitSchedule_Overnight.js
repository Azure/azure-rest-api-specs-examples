const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an Agent resource, which references a hybrid compute machine that can run jobs.
 *
 * @summary creates or updates an Agent resource, which references a hybrid compute machine that can run jobs.
 * x-ms-original-file: 2025-07-01/Agents_CreateOrUpdate_UploadLimitSchedule_Overnight.json
 */
async function agentsCreateOrUpdateWithOvernightUploadLimitSchedule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.agents.createOrUpdate(
    "examples-rg",
    "examples-storageMoverName",
    "examples-agentName",
    {
      properties: {
        arcResourceId:
          "/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName",
        arcVmUuid: "3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9",
        uploadLimitSchedule: {
          weeklyRecurrences: [
            {
              days: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"],
              endTime: { hour: 24, minute: 0 },
              limitInMbps: 2000,
              startTime: { hour: 18, minute: 0 },
            },
            {
              days: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"],
              endTime: { hour: 9, minute: 0 },
              limitInMbps: 2000,
              startTime: { hour: 0, minute: 0 },
            },
          ],
        },
      },
    },
  );
  console.log(result);
}
