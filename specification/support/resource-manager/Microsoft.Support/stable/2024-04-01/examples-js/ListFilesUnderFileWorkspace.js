const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the Files information under a workspace for an Azure subscription.
 *
 * @summary Lists all the Files information under a workspace for an Azure subscription.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/ListFilesUnderFileWorkspace.json
 */
async function listFilesUnderAWorkspace() {
  const fileWorkspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const resArray = new Array();
  for await (let item of client.filesNoSubscription.list(fileWorkspaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
