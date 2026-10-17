const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a ClusterJob
 *
 * @summary get a ClusterJob
 * x-ms-original-file: 2026-10-01/ClusterJobs_Get_AddServerJob.json
 */
async function clusterJobsGetAddServerJob() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.clusterJobs.get("test-rg", "myCluster", "addserver");
  console.log(result);
}
