const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of virtual machine image offers for the specified location, edge zone and publisher.
 *
 * @summary Gets a list of virtual machine image offers for the specified location, edge zone and publisher.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListOffers_MinimumSet_Gen.json
 */
async function virtualMachineImagesEdgeZoneListOffersMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaaaaaaaaa";
  const edgeZone = "aaaaaaaaaaaaaaaaaaaaa";
  const publisherName = "aaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImagesEdgeZone.listOffers(
    location,
    edgeZone,
    publisherName
  );
  console.log(result);
}

virtualMachineImagesEdgeZoneListOffersMinimumSetGen().catch(console.error);
