const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation. If your intent is to create a new compute, do a GET first to verify that it does not exist yet.
 *
 * @summary creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation. If your intent is to create a new compute, do a GET first to verify that it does not exist yet.
 * x-ms-original-file: 2026-03-15-preview/Compute/createOrUpdate/ComputeInstance.json
 */
async function createAnComputeInstanceCompute() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.computeOperations.createOrUpdate(
    "testrg123",
    "workspaces123",
    "compute123",
    {
      location: "eastus",
      properties: {
        computeType: "ComputeInstance",
        properties: {
          applicationSharingPolicy: "Personal",
          autologgerSettings: { mlflowAutologger: "Enabled" },
          computeInstanceAuthorizationType: "personal",
          customServices: [
            {
              name: "rstudio-workbench",
              docker: { privileged: true },
              endpoints: [
                {
                  name: "connect",
                  hostIp: undefined,
                  published: 4444,
                  target: 8787,
                  protocol: "http",
                },
              ],
              environmentVariables: {
                RSP_LICENSE: { type: "local", value: "XXXX-XXXX-XXXX-XXXX-XXXX-XXXX-XXXX" },
              },
              image: { type: "docker", reference: "ghcr.io/azure/rstudio-workbench:latest" },
              kernel: {
                argv: ["option1", "option2", "option3"],
                displayName: "TestKernel",
                language: "python",
              },
              volumes: [
                {
                  type: "bind",
                  readOnly: true,
                  source: "/mnt/azureuser/",
                  target: "/home/testuser/",
                },
              ],
            },
          ],
          enableOSPatching: true,
          enableRootAccess: true,
          enableSSO: true,
          personalComputeInstanceSettings: {
            assignedUser: {
              objectId: "00000000-0000-0000-0000-000000000000",
              tenantId: "00000000-0000-0000-0000-000000000000",
            },
          },
          releaseQuotaOnStop: true,
          sshSettings: { sshPublicAccess: "Disabled" },
          subnet: { id: "test-subnet-resource-id" },
          vmSize: "STANDARD_NC6",
        },
      },
    },
  );
  console.log(result);
}
