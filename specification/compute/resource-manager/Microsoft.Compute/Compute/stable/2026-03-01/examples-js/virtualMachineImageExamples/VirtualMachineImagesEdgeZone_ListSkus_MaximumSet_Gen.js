const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of virtual machine image SKUs for the specified location, edge zone, publisher, and offer.
 *
 * @summary gets a list of virtual machine image SKUs for the specified location, edge zone, publisher, and offer.
 * x-ms-original-file: 2026-03-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListSkus_MaximumSet_Gen.json
 */
async function virtualMachineImagesEdgeZoneListSkusMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImagesEdgeZone.listSkus(
    "aaaaaaaaaaaa",
    "aaaaa",
    "aaaaaaaaaaaa",
    "aaaaaaaaaaaa",
  );
  console.log(result);
}
