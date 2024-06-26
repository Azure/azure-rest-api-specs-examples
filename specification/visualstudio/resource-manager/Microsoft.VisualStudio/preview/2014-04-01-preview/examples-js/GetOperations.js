const { VisualStudioResourceProviderClient } = require("@azure/arm-visualstudio");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of all operations possible on the Microsoft.VisualStudio resource provider.
 *
 * @summary Gets the details of all operations possible on the Microsoft.VisualStudio resource provider.
 * x-ms-original-file: specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/GetOperations.json
 */
async function getAListOfOperationsForThisResourceProvider() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new VisualStudioResourceProviderClient(credential, subscriptionId);
  const result = await client.operations.list();
  console.log(result);
}

getAListOfOperationsForThisResourceProvider().catch(console.error);
