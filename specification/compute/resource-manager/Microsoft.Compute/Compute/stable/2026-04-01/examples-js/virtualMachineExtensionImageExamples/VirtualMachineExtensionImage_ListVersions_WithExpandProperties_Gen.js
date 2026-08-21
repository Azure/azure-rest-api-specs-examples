const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of virtual machine extension image versions.
 *
 * @summary gets a list of virtual machine extension image versions.
 * x-ms-original-file: 2026-04-01/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_ListVersions_WithExpandProperties_Gen.json
 */
async function virtualMachineExtensionImageListVersionsWithExpandPropertiesGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensionImages.listVersions(
    "aaaaaaaaaaaaaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaa",
    { filter: "aaaaaaaaaaaaaaaaaaaaaaaaa", top: 22, orderby: "a", expand: "properties" },
  );
  console.log(result);
}
