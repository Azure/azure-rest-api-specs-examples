const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation. If your intent is to create a new compute, do a GET first to verify that it does not exist yet.
 *
 * @summary Creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation. If your intent is to create a new compute, do a GET first to verify that it does not exist yet.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/Compute/createOrUpdate/KubernetesCompute.json
 */
async function attachAKubernetesCompute() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const workspaceName = "workspaces123";
  const computeName = "compute123";
  const parameters = {
    location: "eastus",
    properties: {
      description: "some compute",
      computeType: "Kubernetes",
      properties: {
        defaultInstanceType: "defaultInstanceType",
        instanceTypes: {
          defaultInstanceType: {
            nodeSelector: {},
            resources: {
              limits: { cpu: "1", memory: "4Gi", "nvidiaCom/gpu": "" },
              requests: { cpu: "1", memory: "4Gi", "nvidiaCom/gpu": "" },
            },
          },
        },
        namespace: "default",
      },
      resourceId:
        "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/Microsoft.ContainerService/managedClusters/compute123-56826-c9b00420020b2",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.computeOperations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    computeName,
    parameters
  );
  console.log(result);
}

attachAKubernetesCompute().catch(console.error);
