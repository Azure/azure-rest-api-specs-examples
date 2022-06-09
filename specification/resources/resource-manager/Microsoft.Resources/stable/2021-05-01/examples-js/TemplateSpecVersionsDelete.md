```javascript
const { TemplateSpecsClient } = require("@azure/arm-templatespecs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a specific version from a Template Spec. When operation completes, status code 200 returned without content.
 *
 * @summary Deletes a specific version from a Template Spec. When operation completes, status code 200 returned without content.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-05-01/examples/TemplateSpecVersionsDelete.json
 */
async function templateSpecVersionsDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "templateSpecRG";
  const templateSpecName = "simpleTemplateSpec";
  const templateSpecVersion = "v1.0";
  const credential = new DefaultAzureCredential();
  const client = new TemplateSpecsClient(credential, subscriptionId);
  const result = await client.templateSpecVersions.delete(
    resourceGroupName,
    templateSpecName,
    templateSpecVersion
  );
  console.log(result);
}

templateSpecVersionsDelete().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-templatespecs_2.0.1/sdk/templatespecs/arm-templatespecs/README.md) on how to add the SDK to your project and authenticate.
