const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the default VM skus resource type
 *
 * @summary Deletes the default VM skus resource type
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/DeleteVmSkus.json
 */
async function deleteVMSkus() {
  const customLocationResourceUri =
    "subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation";
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential);
  const result = await client.beginDeleteVMSkusAndWait(customLocationResourceUri);
  console.log(result);
}
