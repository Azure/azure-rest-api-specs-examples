const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a workbench associated with the project.
 *
 * @summary updates a workbench associated with the project.
 * x-ms-original-file: 2026-09-15-preview/UpdateWorkbench.json
 */
async function updateWorkbench() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.workbenches.update(
    "rgcognitiveservices",
    "myAccount",
    "myProject",
    "myWorkbench",
    { properties: { idleTimeBeforeShutdown: "PT1H" } },
  );
  console.log(result);
}
