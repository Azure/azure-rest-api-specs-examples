const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches the virtual network resource
 *
 * @summary Patches the virtual network resource
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/UpdateVirtualNetwork.json
 */
async function updateVirtualNetwork() {
  const subscriptionId =
    process.env["HYBRIDCONTAINERSERVICE_SUBSCRIPTION_ID"] || "a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b";
  const resourceGroupName =
    process.env["HYBRIDCONTAINERSERVICE_RESOURCE_GROUP"] || "test-arcappliance-resgrp";
  const virtualNetworkName = "test-vnet-static";
  const virtualNetworks = {
    tags: { additionalProperties: "sample" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential, subscriptionId);
  const result = await client.virtualNetworks.beginUpdateAndWait(
    resourceGroupName,
    virtualNetworkName,
    virtualNetworks,
  );
  console.log(result);
}
