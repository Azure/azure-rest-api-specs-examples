const { ImageBuilderClient } = require("@azure/arm-imagebuilder");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the tags for this Virtual Machine Image Template
 *
 * @summary Update the tags for this Virtual Machine Image Template
 * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/UpdateImageTemplateTags.json
 */
async function updateTheTagsForAnImageTemplate() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const imageTemplateName = "myImageTemplate";
  const parameters = {
    tags: { newTag: "new-value" },
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

updateTheTagsForAnImageTemplate().catch(console.error);
