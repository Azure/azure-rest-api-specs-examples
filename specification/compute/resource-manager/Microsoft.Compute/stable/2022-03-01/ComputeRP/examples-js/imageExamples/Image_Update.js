const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an image.
 *
 * @summary Update an image.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/imageExamples/Image_Update.json
 */
async function updatesTagsOfAnImage() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const imageName = "myImage";
  const parameters = {
    hyperVGeneration: "V1",
    sourceVirtualMachine: {
      id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM",
    },
    tags: { department: "HR" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.images.beginUpdateAndWait(resourceGroupName, imageName, parameters);
  console.log(result);
}

updatesTagsOfAnImage().catch(console.error);
