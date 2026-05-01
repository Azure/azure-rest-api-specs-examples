const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to execute listWithProperties
 *
 * @summary execute listWithProperties
 * x-ms-original-file: 2025-11-01/virtualMachineImageExamples/VirtualMachineImages_ListWithProperties_MinimumSet_Gen.json
 */
async function virtualMachineImagesListWithPropertiesMinimumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.listWithProperties(
    "eastus",
    "MicrosoftWindowsServer",
    "WindowsServer",
    "2022-datacenter-azure-edition",
    "Properties",
  );
  console.log(result);
}
