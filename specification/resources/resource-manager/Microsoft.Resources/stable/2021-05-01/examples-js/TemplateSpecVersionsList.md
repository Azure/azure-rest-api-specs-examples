Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-templatespecs_2.0.1/sdk/templatespecs/arm-templatespecs/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { TemplateSpecsClient } = require("@azure/arm-templatespecs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the Template Spec versions in the specified Template Spec.
 *
 * @summary Lists all the Template Spec versions in the specified Template Spec.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-05-01/examples/TemplateSpecVersionsList.json
 */
async function templateSpecVersionsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "templateSpecRG";
  const templateSpecName = "simpleTemplateSpec";
  const credential = new DefaultAzureCredential();
  const client = new TemplateSpecsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.templateSpecVersions.list(resourceGroupName, templateSpecName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

templateSpecVersionsList().catch(console.error);
```
