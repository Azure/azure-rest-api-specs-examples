const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new file workspace.
 *
 * @summary Creates a new file workspace.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/CreateFileWorkspace.json
 */
async function createAFileWorkspace() {
  const fileWorkspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const result = await client.fileWorkspacesNoSubscription.create(fileWorkspaceName);
  console.log(result);
}
