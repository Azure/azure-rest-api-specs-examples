const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details for a specific file workspace.
 *
 * @summary Gets details for a specific file workspace.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/GetFileWorkspaceDetails.json
 */
async function getDetailsOfAFileWorkspace() {
  const fileWorkspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const result = await client.fileWorkspacesNoSubscription.get(fileWorkspaceName);
  console.log(result);
}
