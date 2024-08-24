const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation. If your intent is to create a new compute, do a GET first to verify that it does not exist yet.
 *
 * @summary Creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation. If your intent is to create a new compute, do a GET first to verify that it does not exist yet.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Compute/createOrUpdate/ComputeInstance.json
 */
async function createAnComputeInstanceCompute() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "testrg123";
  const workspaceName = "workspaces123";
  const computeName = "compute123";
  const parameters = {
    location: "eastus",
    properties: {
      computeType: "ComputeInstance",
      properties: {
        applicationSharingPolicy: "Personal",
        computeInstanceAuthorizationType: "personal",
        customServices: [
          {
            name: "rstudio",
            docker: { privileged: true },
            endpoints: [
              {
                name: "connect",
                published: 8787,
                target: 8787,
                protocol: "http",
              },
            ],
            environmentVariables: {
              testVariable: { type: "local", value: "test_value" },
            },
            image: {
              type: "docker",
              reference: "ghcr.io/azure/rocker-rstudio-ml-verse:latest",
            },
            volumes: [
              {
                type: "bind",
                readOnly: false,
                source: "/home/azureuser/cloudfiles",
                target: "/home/azureuser/cloudfiles",
              },
            ],
          },
        ],
        personalComputeInstanceSettings: {
          assignedUser: {
            objectId: "00000000-0000-0000-0000-000000000000",
            tenantId: "00000000-0000-0000-0000-000000000000",
          },
        },
        sshSettings: { sshPublicAccess: "Disabled" },
        subnet: { id: "test-subnet-resource-id" },
        vmSize: "STANDARD_NC6",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.computeOperations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    computeName,
    parameters,
  );
  console.log(result);
}
