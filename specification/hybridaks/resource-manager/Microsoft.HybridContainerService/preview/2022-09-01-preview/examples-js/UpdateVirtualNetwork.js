const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches the Hybrid AKS virtual network
 *
 * @summary Patches the Hybrid AKS virtual network
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/UpdateVirtualNetwork.json
 */
async function updateVirtualNetwork() {
  const subscriptionId =
    process.env["HYBRIDCONTAINERSERVICE_SUBSCRIPTION_ID"] || "a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b";
  const resourceGroupName =
    process.env["HYBRIDCONTAINERSERVICE_RESOURCE_GROUP"] || "test-arcappliance-resgrp";
  const virtualNetworksName = "test-vnet-static";
  const virtualNetworks = {
    tags: { additionalProperties: "sample" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential, subscriptionId);
  const result = await client.virtualNetworksOperations.beginUpdateAndWait(
    resourceGroupName,
    virtualNetworksName,
    virtualNetworks
  );
  console.log(result);
}
