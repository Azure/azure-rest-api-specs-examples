const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of virtual machine image publishers for the specified Azure location and edge zone.
 *
 * @summary gets a list of virtual machine image publishers for the specified Azure location and edge zone.
 * x-ms-original-file: 2025-11-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListPublishers_MaximumSet_Gen.json
 */
async function virtualMachineImagesEdgeZoneListPublishersMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImagesEdgeZone.listPublishers(
    "aaaaaa",
    "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
  );
  console.log(result);
}
