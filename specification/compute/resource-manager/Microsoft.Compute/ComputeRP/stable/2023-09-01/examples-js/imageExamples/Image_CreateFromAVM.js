const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an image.
 *
 * @summary Create or update an image.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/imageExamples/Image_CreateFromAVM.json
 */
async function createAVirtualMachineImageFromAnExistingVirtualMachine() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const imageName = "myImage";
  const parameters = {
    location: "West US",
    sourceVirtualMachine: {
      id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.images.beginCreateOrUpdateAndWait(
    resourceGroupName,
    imageName,
    parameters,
  );
  console.log(result);
}
