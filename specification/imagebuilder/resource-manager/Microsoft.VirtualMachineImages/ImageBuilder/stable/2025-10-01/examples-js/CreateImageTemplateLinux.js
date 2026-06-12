const { ImageBuilderClient } = require("@azure/arm-imagebuilder");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a virtual machine image template
 *
 * @summary create or update a virtual machine image template
 * x-ms-original-file: 2025-10-01/CreateImageTemplateLinux.json
 */
async function createAnImageTemplateForLinux() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ImageBuilderClient(credential, subscriptionId);
  const result = await client.virtualMachineImageTemplates.createOrUpdate(
    "myResourceGroup",
    "myImageTemplate",
    {
      identity: {
        type: "UserAssigned",
        userAssignedIdentities: {
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity_1":
            {},
        },
      },
      location: "westus",
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
      source: {
        type: "ManagedImage",
        imageId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/images/source_image",
      },
      vmProfile: {
        osDiskSizeGB: 64,
        vmSize: "Standard_D2s_v3",
        vnetConfig: {
          subnetId:
            "/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet_name/subnets/subnet_name",
        },
      },
      tags: { imagetemplate_tag1: "IT_T1", imagetemplate_tag2: "IT_T2" },
    },
  );
  console.log(result);
}
