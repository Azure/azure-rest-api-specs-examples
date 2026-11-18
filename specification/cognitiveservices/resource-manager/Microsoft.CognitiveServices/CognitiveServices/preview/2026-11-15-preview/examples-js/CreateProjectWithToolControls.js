const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create Cognitive Services Account's Project. Project is a sub-resource of an account which give AI developer it's individual container to work on.
 *
 * @summary create Cognitive Services Account's Project. Project is a sub-resource of an account which give AI developer it's individual container to work on.
 * x-ms-original-file: 2026-11-15-preview/CreateProjectWithToolControls.json
 */
async function createAFoundryProjectWithToolControls() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.projects.create("myResourceGroup", "myAccount", "myProject", {
    location: "westus",
    properties: {
      displayName: "My project",
      toolControls: {
        defaultAction: "Deny",
        allowedTools: [
          { type: "code_interpreter" },
          { type: "mcp", serverUrl: "https://*.contoso.com/mcp" },
        ],
      },
    },
  });
  console.log(result);
}
