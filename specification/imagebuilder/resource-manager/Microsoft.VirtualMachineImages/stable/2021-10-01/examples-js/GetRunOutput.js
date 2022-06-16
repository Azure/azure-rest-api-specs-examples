const { ImageBuilderClient } = require("@azure/arm-imagebuilder");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified run output for the specified image template resource
 *
 * @summary Get the specified run output for the specified image template resource
 * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/GetRunOutput.json
 */
async function retrieveSingleRunOutput() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const imageTemplateName = "myImageTemplate";
  const runOutputName = "myManagedImageOutput";
  const credential = new DefaultAzureCredential();
  const client = new ImageBuilderClient(credential, subscriptionId);
  const result = await client.virtualMachineImageTemplates.getRunOutput(
    resourceGroupName,
    imageTemplateName,
    runOutputName
  );
  console.log(result);
}

retrieveSingleRunOutput().catch(console.error);
