const { ImageBuilderClient } = require("@azure/arm-imagebuilder");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create artifacts from a existing image template
 *
 * @summary Create artifacts from a existing image template
 * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/RunImageTemplate.json
 */
async function createImageSFromExistingImageTemplate() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const imageTemplateName = "myImageTemplate";
  const credential = new DefaultAzureCredential();
  const client = new ImageBuilderClient(credential, subscriptionId);
  const result = await client.virtualMachineImageTemplates.beginRunAndWait(
    resourceGroupName,
    imageTemplateName
  );
  console.log(result);
}

createImageSFromExistingImageTemplate().catch(console.error);
