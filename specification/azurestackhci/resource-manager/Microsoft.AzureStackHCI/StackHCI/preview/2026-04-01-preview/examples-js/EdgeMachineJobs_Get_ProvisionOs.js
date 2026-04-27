const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a EdgeMachineJob
 *
 * @summary get a EdgeMachineJob
 * x-ms-original-file: 2026-04-01-preview/EdgeMachineJobs_Get_ProvisionOs.json
 */
async function edgeMachineJobsGetProvisionOs() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.edgeMachineJobs.get("ArcInstance-rg", "machine1", "ProvisionOs");
  console.log(result);
}
