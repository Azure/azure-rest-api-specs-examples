const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a virtual machine image in an edge zone.
 *
 * @summary gets a virtual machine image in an edge zone.
 * x-ms-original-file: 2025-11-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_Get_MaximumSet_Gen.json
 */
async function virtualMachineImagesEdgeZoneGetMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImagesEdgeZone.get(
    "aaaaaaaaaaaaaaaaaaaaaaaa",
    "aaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
    "aaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
  );
  console.log(result);
}
