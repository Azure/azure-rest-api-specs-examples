Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-templatespecs_2.0.1/sdk/templatespecs/arm-templatespecs/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { TemplateSpecsClient } = require("@azure/arm-templatespecs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Template Spec version.
 *
 * @summary Creates or updates a Template Spec version.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-05-01/examples/TemplateSpecVersionsCreate.json
 */
async function templateSpecVersionsCreateUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "templateSpecRG";
  const templateSpecName = "simpleTemplateSpec";
  const templateSpecVersion = "v1.0";
  const templateSpecVersionModel = {
    description: "This is version v1.0 of our template content",
    location: "eastus",
    mainTemplate: {
      $schema: "http://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
      contentVersion: "1.0.0.0",
      parameters: {},
      resources: [],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new TemplateSpecsClient(credential, subscriptionId);
  const result = await client.templateSpecVersions.createOrUpdate(
    resourceGroupName,
    templateSpecName,
    templateSpecVersion,
    templateSpecVersionModel
  );
  console.log(result);
}

templateSpecVersionsCreateUpdate().catch(console.error);
```
