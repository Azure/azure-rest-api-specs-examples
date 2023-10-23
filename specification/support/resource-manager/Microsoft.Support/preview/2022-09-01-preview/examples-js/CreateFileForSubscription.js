const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new file under a workspace for the specified subscription.
 *
 * @summary Creates a new file under a workspace for the specified subscription.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/CreateFileForSubscription.json
 */
async function createAFileWorkspace() {
  const subscriptionId = process.env["SUPPORT_SUBSCRIPTION_ID"] || "subid";
  const fileWorkspaceName = "testworkspace";
  const fileName = "test.txt";
  const createFileParameters = {
    chunkSize: 41423,
    fileSize: 41423,
    numberOfChunks: 1,
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.files.create(fileWorkspaceName, fileName, createFileParameters);
  console.log(result);
}
