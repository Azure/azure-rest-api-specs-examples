const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a virtual machine image in an edge zone.
 *
 * @summary Gets a virtual machine image in an edge zone.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_Get_MaximumSet_Gen.json
 */
async function virtualMachineImagesEdgeZoneGetMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const location = "aaaaaaaaaaaaaaaaaaaaaaaa";
  const edgeZone = "aaaaaaaa";
  const publisherName = "aaaaaaaaaaaaaaaaaaaaaaa";
  const offer = "aaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const skus = "aaaaaaaaaa";
  const version = "aaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImagesEdgeZone.get(
    location,
    edgeZone,
    publisherName,
    offer,
    skus,
    version,
  );
  console.log(result);
}
