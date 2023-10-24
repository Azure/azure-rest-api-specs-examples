const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns details of a specific file in a work space.
 *
 * @summary Returns details of a specific file in a work space.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/GetFileDetails.json
 */
async function getDetailsOfASubscriptionFile() {
  const fileWorkspaceName = "testworkspace";
  const fileName = "test.txt";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const result = await client.filesNoSubscription.get(fileWorkspaceName, fileName);
  console.log(result);
}
