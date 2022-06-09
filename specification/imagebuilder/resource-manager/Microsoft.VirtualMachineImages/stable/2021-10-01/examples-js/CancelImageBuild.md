```javascript
const { ImageBuilderClient } = require("@azure/arm-imagebuilder");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Cancel the long running image build based on the image template
 *
 * @summary Cancel the long running image build based on the image template
 * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/CancelImageBuild.json
 */
async function cancelTheImageBuildBasedOnTheImageTemplate() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const imageTemplateName = "myImageTemplate";
  const credential = new DefaultAzureCredential();
  const client = new ImageBuilderClient(credential, subscriptionId);
  const result = await client.virtualMachineImageTemplates.beginCancelAndWait(
    resourceGroupName,
    imageTemplateName
  );
  console.log(result);
}

cancelTheImageBuildBasedOnTheImageTemplate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-imagebuilder_1.0.2/sdk/imagebuilder/arm-imagebuilder/README.md) on how to add the SDK to your project and authenticate.
