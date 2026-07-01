const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a compute associated with the Cognitive Services account.
 *
 * @summary creates or updates a compute associated with the Cognitive Services account.
 * x-ms-original-file: 2026-05-15-preview/PutContainerInstanceCompute.json
 */
async function putContainerInstanceCompute() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.computes.createOrUpdate(
    "rgcognitiveservices",
    "myAccount",
    "myContainerInstance",
    {
      properties: {
        computeType: "ContainerInstance",
        targetClusterId:
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/computes/myCluster",
        imageLink: "mcr.microsoft.com/azureml/curated/pytorch-gpu:latest",
        idleTimeBeforeShutdown: "PT30M",
        sshSettings: {
          sshPublicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQ...",
          adminEnabled: true,
        },
      },
      location: "eastus",
      identity: {
        type: "UserAssigned",
        userAssignedIdentities: {
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity":
            {},
        },
      },
    },
  );
  console.log(result);
}
