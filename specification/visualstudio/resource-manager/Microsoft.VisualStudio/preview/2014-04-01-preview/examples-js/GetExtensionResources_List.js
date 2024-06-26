const { VisualStudioResourceProviderClient } = require("@azure/arm-visualstudio");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of the extension resources created within the resource group.
 *
 * @summary Gets the details of the extension resources created within the resource group.
 * x-ms-original-file: specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/GetExtensionResources_List.json
 */
async function getAListOfExtensionResourcesWithinTheResourceGroup() {
  const subscriptionId = "0de7f055-dbea-498d-8e9e-da287eedca90";
  const resourceGroupName = "VS-Example-Group";
  const accountResourceName = "ExampleAccount";
  const credential = new DefaultAzureCredential();
  const client = new VisualStudioResourceProviderClient(credential, subscriptionId);
  const result = await client.extensions.listByAccount(resourceGroupName, accountResourceName);
  console.log(result);
}

getAListOfExtensionResourcesWithinTheResourceGroup().catch(console.error);
