const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to
 *
 * @summary
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/virtualMachineImageExamples/VirtualMachineImages_ListWithProperties_MinimumSet_Gen.json
 */
async function virtualMachineImagesListWithPropertiesMinimumSet() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const location = "eastus";
  const publisherName = "MicrosoftWindowsServer";
  const offer = "WindowsServer";
  const skus = "2022-datacenter-azure-edition";
  const expand = "Properties";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.virtualMachineImages.listWithProperties(
    location,
    publisherName,
    offer,
    skus,
    expand,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
