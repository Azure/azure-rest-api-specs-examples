const { VisualStudioResourceProviderClient } = require("@azure/arm-visualstudio");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all Visual Studio Team Services account resources under the resource group linked to the specified Azure subscription.
 *
 * @summary Gets all Visual Studio Team Services account resources under the resource group linked to the specified Azure subscription.
 * x-ms-original-file: specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/GetResources_List.json
 */
async function getAListOfAccountResourcesInTheResourceGroup() {
  const subscriptionId = "0de7f055-dbea-498d-8e9e-da287eedca90";
  const resourceGroupName = "VS-Example-Group";
  const credential = new DefaultAzureCredential();
  const client = new VisualStudioResourceProviderClient(credential, subscriptionId);
  const result = await client.accounts.listByResourceGroup(resourceGroupName);
  console.log(result);
}

getAListOfAccountResourcesInTheResourceGroup().catch(console.error);
