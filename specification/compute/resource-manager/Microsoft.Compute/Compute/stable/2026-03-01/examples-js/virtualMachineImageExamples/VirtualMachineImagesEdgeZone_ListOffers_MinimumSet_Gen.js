const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of virtual machine image offers for the specified location, edge zone and publisher.
 *
 * @summary gets a list of virtual machine image offers for the specified location, edge zone and publisher.
 * x-ms-original-file: 2026-03-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListOffers_MinimumSet_Gen.json
 */
async function virtualMachineImagesEdgeZoneListOffersMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImagesEdgeZone.listOffers(
    "aaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
  );
  console.log(result);
}
