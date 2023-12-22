const { ImageBuilderClient } = require("@azure/arm-imagebuilder");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a virtual machine image template
 *
 * @summary Create or update a virtual machine image template
 * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2023-07-01/examples/CreateImageTemplateLinux.json
 */
async function createAnImageTemplateForLinux() {
  const subscriptionId = process.env["IMAGEBUILDER_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["IMAGEBUILDER_RESOURCE_GROUP"] || "myResourceGroup";
  const imageTemplateName = "myImageTemplate";
  const parameters = {
    customize: [
      {
        name: "Shell Customizer Example",
        type: "Shell",
        scriptUri: "https://example.com/path/to/script.sh",
      },
    ],
    distribute: [
      {
        type: "ManagedImage",
        artifactTags: { tagName: "value" },
        imageId:
          "/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Compute/images/image_it_1",
        location: "1_location",
        runOutputName: "image_it_pir_1",
      },
    ],
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000000000000000000000000000/resourcegroups/rg1/providers/MicrosoftManagedIdentity/userAssignedIdentities/identity1":
          {},
      },
    },
    location: "westus",
    source: {
      type: "ManagedImage",
      imageId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/images/source_image",
    },
    tags: { imagetemplateTag1: "IT_T1", imagetemplateTag2: "IT_T2" },
    vmProfile: {
      osDiskSizeGB: 64,
      vmSize: "Standard_D2s_v3",
      vnetConfig: {
        subnetId:
          "/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet_name/subnets/subnet_name",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ImageBuilderClient(credential, subscriptionId);
  const result = await client.virtualMachineImageTemplates.beginCreateOrUpdateAndWait(
    resourceGroupName,
    imageTemplateName,
    parameters,
  );
  console.log(result);
}
