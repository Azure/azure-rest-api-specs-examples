const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of virtual machine image offers for the specified location, edge zone and publisher.
 *
 * @summary Gets a list of virtual machine image offers for the specified location, edge zone and publisher.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListOffers_MaximumSet_Gen.json
 */
async function virtualMachineImagesEdgeZoneListOffersMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const location = "aaaaaaaaaaaaaaaaaa";
  const edgeZone = "aaaaaaaaaaaaaaaaaaa";
  const publisherName = "aaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImagesEdgeZone.listOffers(
    location,
    edgeZone,
    publisherName
  );
  console.log(result);
}
