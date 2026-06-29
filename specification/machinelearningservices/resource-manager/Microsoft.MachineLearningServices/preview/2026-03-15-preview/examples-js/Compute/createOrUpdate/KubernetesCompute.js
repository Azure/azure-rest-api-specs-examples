const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation. If your intent is to create a new compute, do a GET first to verify that it does not exist yet.
 *
 * @summary creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation. If your intent is to create a new compute, do a GET first to verify that it does not exist yet.
 * x-ms-original-file: 2026-03-15-preview/Compute/createOrUpdate/KubernetesCompute.json
 */
async function attachAKubernetesCompute() {
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
        description: "some compute",
        computeType: "Kubernetes",
        properties: {
          defaultInstanceType: "defaultInstanceType",
          instanceTypes: {
            defaultInstanceType: {
              nodeSelector: {},
              resources: {
                limits: { cpu: "1", memory: "4Gi" },
                requests: { cpu: "1", memory: "4Gi" },
              },
            },
          },
          namespace: "default",
        },
        resourceId:
          "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/Microsoft.ContainerService/managedClusters/compute123-56826-c9b00420020b2",
      },
    },
  );
  console.log(result);
}
