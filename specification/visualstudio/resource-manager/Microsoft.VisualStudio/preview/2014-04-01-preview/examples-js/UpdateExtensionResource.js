const { VisualStudioResourceProviderClient } = require("@azure/arm-visualstudio");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing extension registration for the Visual Studio Team Services account.
 *
 * @summary Updates an existing extension registration for the Visual Studio Team Services account.
 * x-ms-original-file: specification/visualstudio/resource-manager/Microsoft.VisualStudio/preview/2014-04-01-preview/examples/UpdateExtensionResource.json
 */
async function updateAnExtensionResource() {
  const subscriptionId = "0de7f055-dbea-498d-8e9e-da287eedca90";
  const resourceGroupName = "VS-Example-Group";
  const accountResourceName = "ExampleAccount";
  const extensionResourceName = "Example";
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
  const result = await client.extensions.update(
    resourceGroupName,
    accountResourceName,
    extensionResourceName,
    body
  );
  console.log(result);
}

updateAnExtensionResource().catch(console.error);
