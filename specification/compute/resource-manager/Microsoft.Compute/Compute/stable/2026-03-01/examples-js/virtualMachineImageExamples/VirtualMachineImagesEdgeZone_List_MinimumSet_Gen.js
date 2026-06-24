const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of all virtual machine image versions for the specified location, edge zone, publisher, offer, and SKU.
 *
 * @summary gets a list of all virtual machine image versions for the specified location, edge zone, publisher, offer, and SKU.
 * x-ms-original-file: 2026-03-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_List_MinimumSet_Gen.json
 */
async function virtualMachineImagesEdgeZoneListMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImagesEdgeZone.list(
    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
    "aaaaa",
    "aaaaaaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
  );
  console.log(result);
}
