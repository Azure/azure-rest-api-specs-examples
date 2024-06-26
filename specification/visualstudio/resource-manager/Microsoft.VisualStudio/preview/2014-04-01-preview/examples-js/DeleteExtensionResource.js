const { VisualStudioResourceProviderClient } = require("@azure/arm-visualstudio");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Removes an extension resource registration for a Visual Studio Team Services account.
 *
 * @summary Removes an extension resource registration for a Visual Studio Team Services account.
 * x-ms-original-file: specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/DeleteExtensionResource.json
 */
async function deleteAnExtensionResource() {
  const subscriptionId = "0de7f055-dbea-498d-8e9e-da287eedca90";
  const resourceGroupName = "VS-Example-Group";
  const accountResourceName = "Example";
  const extensionResourceName = "ms.example";
  const credential = new DefaultAzureCredential();
  const client = new VisualStudioResourceProviderClient(credential, subscriptionId);
  const result = await client.extensions.delete(
    resourceGroupName,
    accountResourceName,
    extensionResourceName
  );
  console.log(result);
}

deleteAnExtensionResource().catch(console.error);
