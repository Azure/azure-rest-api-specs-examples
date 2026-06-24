const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of all virtual machine image versions for the specified location, edge zone, publisher, offer, and SKU.
 *
 * @summary gets a list of all virtual machine image versions for the specified location, edge zone, publisher, offer, and SKU.
 * x-ms-original-file: 2026-03-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_List_MaximumSet_Gen.json
 */
async function virtualMachineImagesEdgeZoneListMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImagesEdgeZone.list(
    "aaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaaaaaa",
    "aaaa",
    "aaaaaaaaaaaaaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaaaa",
    { expand: "aaaaaaaaaaaaaaaaaaaaaaaa", top: 12, orderby: "aaaaaaaaaaaaaaaaaaaaaaaaaaaa" },
  );
  console.log(result);
}
