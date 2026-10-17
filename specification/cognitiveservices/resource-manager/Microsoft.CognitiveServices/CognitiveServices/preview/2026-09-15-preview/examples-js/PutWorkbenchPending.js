const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a workbench associated with the project.
 *
 * @summary creates or updates a workbench associated with the project.
 * x-ms-original-file: 2026-09-15-preview/PutWorkbenchPending.json
 */
async function putWorkbenchPending() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.workbenches.createOrUpdate(
    "rgcognitiveservices",
    "myAccount",
    "myProject",
    "myWorkbench",
    {
      properties: {
        targetClusterId:
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/computes/myCluster",
        imageLink: "mcr.microsoft.com/azureml/curated/pytorch-gpu:latest",
      },
    },
  );
  console.log(result);
}
