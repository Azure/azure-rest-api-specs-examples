Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-imagebuilder_1.0.2/sdk/imagebuilder/arm-imagebuilder/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ImageBuilderClient } = require("@azure/arm-imagebuilder");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a virtual machine image template
 *
 * @summary Delete a virtual machine image template
 * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/DeleteImageTemplate.json
 */
async function deleteAnImageTemplate() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const imageTemplateName = "myImageTemplate";
  const credential = new DefaultAzureCredential();
  const client = new ImageBuilderClient(credential, subscriptionId);
  const result = await client.virtualMachineImageTemplates.beginDeleteAndWait(
    resourceGroupName,
    imageTemplateName
  );
  console.log(result);
}

deleteAnImageTemplate().catch(console.error);
```
