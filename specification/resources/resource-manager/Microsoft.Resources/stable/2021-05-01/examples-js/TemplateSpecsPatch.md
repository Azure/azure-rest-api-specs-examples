```javascript
const { TemplateSpecsClient } = require("@azure/arm-templatespecs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates Template Spec tags with specified values.
 *
 * @summary Updates Template Spec tags with specified values.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-05-01/examples/TemplateSpecsPatch.json
 */
async function templateSpecsPatch() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "templateSpecRG";
  const templateSpecName = "simpleTemplateSpec";
  const templateSpec = { tags: { myTag: "My Value" } };
  const options = { templateSpec };
  const credential = new DefaultAzureCredential();
  const client = new TemplateSpecsClient(credential, subscriptionId);
  const result = await client.templateSpecs.update(resourceGroupName, templateSpecName, options);
  console.log(result);
}

templateSpecsPatch().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-templatespecs_2.0.1/sdk/templatespecs/arm-templatespecs/README.md) on how to add the SDK to your project and authenticate.
