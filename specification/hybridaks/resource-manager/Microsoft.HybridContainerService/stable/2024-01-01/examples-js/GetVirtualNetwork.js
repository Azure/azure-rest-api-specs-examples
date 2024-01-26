const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified virtual network resource
 *
 * @summary Gets the specified virtual network resource
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/GetVirtualNetwork.json
 */
async function getVirtualNetwork() {
  const subscriptionId =
    process.env["HYBRIDCONTAINERSERVICE_SUBSCRIPTION_ID"] || "a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b";
  const resourceGroupName =
    process.env["HYBRIDCONTAINERSERVICE_RESOURCE_GROUP"] || "test-arcappliance-resgrp";
  const virtualNetworkName = "test-vnet-static";
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential, subscriptionId);
  const result = await client.virtualNetworks.retrieve(resourceGroupName, virtualNetworkName);
  console.log(result);
}
