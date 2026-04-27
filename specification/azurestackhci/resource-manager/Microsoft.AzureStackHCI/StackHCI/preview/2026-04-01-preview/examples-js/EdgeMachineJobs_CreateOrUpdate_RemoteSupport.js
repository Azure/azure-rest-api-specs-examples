const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a EdgeMachineJob
 *
 * @summary create a EdgeMachineJob
 * x-ms-original-file: 2026-04-01-preview/EdgeMachineJobs_CreateOrUpdate_RemoteSupport.json
 */
async function edgeMachineJobsCreateOrUpdateRemoteSupport() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.edgeMachineJobs.createOrUpdate(
    "ArcInstance-rg",
    "machine1",
    "RemoteSupport",
    {
      properties: {
        jobType: "RemoteSupport",
        accessLevel: "Diagnostics",
        type: "Enable",
        expirationTimestamp: new Date("2024-01-29T10:43:27.9471574Z"),
      },
    },
  );
  console.log(result);
}
