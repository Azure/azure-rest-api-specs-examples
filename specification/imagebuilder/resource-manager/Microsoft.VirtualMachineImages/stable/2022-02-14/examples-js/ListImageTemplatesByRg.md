Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-imagebuilder_2.0.0/sdk/imagebuilder/arm-imagebuilder/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ImageBuilderClient } = require("@azure/arm-imagebuilder");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the VM image templates associated with the specified resource group.
 *
 * @summary Gets information about the VM image templates associated with the specified resource group.
 * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/ListImageTemplatesByRg.json
 */
async function listImagesByResourceGroup() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new ImageBuilderClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachineImageTemplates.listByResourceGroup(
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listImagesByResourceGroup().catch(console.error);
```
