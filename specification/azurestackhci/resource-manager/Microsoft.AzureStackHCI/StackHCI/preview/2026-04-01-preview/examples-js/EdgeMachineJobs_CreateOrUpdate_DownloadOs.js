const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a EdgeMachineJob
 *
 * @summary create a EdgeMachineJob
 * x-ms-original-file: 2026-04-01-preview/EdgeMachineJobs_CreateOrUpdate_DownloadOs.json
 */
async function edgeMachineJobsCreateOrUpdateDownloadOs() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.edgeMachineJobs.createOrUpdate(
    "ArcInstance-rg",
    "machine1",
    "DownloadOs",
    {
      properties: {
        jobType: "DownloadOs",
        deploymentMode: "Deploy",
        downloadRequest: {
          target: "AzureLinux",
          osProfile: {
            osName: "AzureLinux",
            osType: "AzureLinux",
            osVersion: "3.0",
            osImageLocation: "https://aka.ms/aep/azlinux3.0",
            vsrVersion: "1.0.0",
            imageHash: "sha256:a8b1c2d3e4f5a6b7c8d9e0f1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1",
            gpgPubKey:
              "LS0tLS1CRUdJTiBQR1AgUFVCTElDIEtFWSBCTE9DSy0tLS0tXG5WZXJzaW9uOiBHbnVQRyB2MlxuXG5tUUVOQkZYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYXG4tLS0tLUVORCBQR1AgUFVCTElDIEtFWSBCTE9DSy0tLS0t",
          },
        },
      },
    },
  );
  console.log(result);
}
