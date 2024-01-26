const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the supported kubernetes versions for the specified custom location
 *
 * @summary Lists the supported kubernetes versions for the specified custom location
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/GetKubernetesVersions.json
 */
async function getKubernetesVersions() {
  const customLocationResourceUri =
    "subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation";
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential);
  const result = await client.getKubernetesVersions(customLocationResourceUri);
  console.log(result);
}
