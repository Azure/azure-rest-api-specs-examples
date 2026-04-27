const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a EdgeMachineJob
 *
 * @summary create a EdgeMachineJob
 * x-ms-original-file: 2026-04-01-preview/EdgeMachineJobs_CreateOrUpdate_CollectLog.json
 */
async function edgeMachineJobsCreateOrUpdateCollectLog() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.edgeMachineJobs.createOrUpdate(
    "ArcInstance-rg",
    "machine1",
    "triggerLogCollection",
    { properties: { jobType: "EdgeMachineJobProperties", deploymentMode: "Validate" } },
  );
  console.log(result);
}
