const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets details for a specific file workspace.
 *
 * @summary gets details for a specific file workspace.
 * x-ms-original-file: 2025-06-01-preview/GetFileWorkspaceDetails.json
 */
async function getDetailsOfAFileWorkspace() {
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const result = await client.fileWorkspacesNoSubscription.get("testworkspace");
  console.log(result);
}
