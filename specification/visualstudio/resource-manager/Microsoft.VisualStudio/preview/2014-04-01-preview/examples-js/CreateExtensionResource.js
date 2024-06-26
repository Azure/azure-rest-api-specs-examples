const { VisualStudioResourceProviderClient } = require("@azure/arm-visualstudio");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Registers the extension with a Visual Studio Team Services account.
 *
 * @summary Registers the extension with a Visual Studio Team Services account.
 * x-ms-original-file: specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/CreateExtensionResource.json
 */
async function createAnExtensionResource() {
  const subscriptionId = "0de7f055-dbea-498d-8e9e-da287eedca90";
  const resourceGroupName = "VS-Example-Group";
  const accountResourceName = "ExampleAccount";
  const extensionResourceName = "ms.example";
  const body = {
    location: "Central US",
    plan: {
      name: "ExamplePlan",
      product: "ExampleExtensionName",
      promotionCode: "",
      publisher: "ExampleExtensionPublisher",
      version: "1.0",
    },
    properties: {},
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new VisualStudioResourceProviderClient(credential, subscriptionId);
  const result = await client.extensions.create(
    resourceGroupName,
    accountResourceName,
    extensionResourceName,
    body
  );
  console.log(result);
}

createAnExtensionResource().catch(console.error);
