Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-imagebuilder_2.0.0/sdk/imagebuilder/arm-imagebuilder/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ImageBuilderClient } = require("@azure/arm-imagebuilder");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the tags for this Virtual Machine Image Template
 *
 * @summary Update the tags for this Virtual Machine Image Template
 * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/UpdateImageTemplateToRemoveIdentities.json
 */
async function removeIdentitiesForAnImageTemplate() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const imageTemplateName = "myImageTemplate";
  const parameters = {
    identity: { type: "None" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ImageBuilderClient(credential, subscriptionId);
  const result = await client.virtualMachineImageTemplates.beginUpdateAndWait(
    resourceGroupName,
    imageTemplateName,
    parameters
  );
  console.log(result);
}

removeIdentitiesForAnImageTemplate().catch(console.error);
```
