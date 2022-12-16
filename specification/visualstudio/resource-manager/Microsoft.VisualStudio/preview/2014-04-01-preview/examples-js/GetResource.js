const { VisualStudioResourceProviderClient } = require("@azure/arm-visualstudio");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Visual Studio Team Services account resource details.
 *
 * @summary Gets the Visual Studio Team Services account resource details.
 * x-ms-original-file: specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/GetResource.json
 */
async function getAnAccountResource() {
  const subscriptionId = "0de7f055-dbea-498d-8e9e-da287eedca90";
  const resourceGroupName = "VS-Example-Group";
  const resourceName = "Example";
  const credential = new DefaultAzureCredential();
  const client = new VisualStudioResourceProviderClient(credential, subscriptionId);
  const result = await client.accounts.get(resourceGroupName, resourceName);
  console.log(result);
}

getAnAccountResource().catch(console.error);
