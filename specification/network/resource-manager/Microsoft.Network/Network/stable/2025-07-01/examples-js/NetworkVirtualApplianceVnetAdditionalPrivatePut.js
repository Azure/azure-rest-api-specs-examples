const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the specified Network Virtual Appliance.
 *
 * @summary creates or updates the specified Network Virtual Appliance.
 * x-ms-original-file: 2025-07-01/NetworkVirtualApplianceVnetAdditionalPrivatePut.json
 */
async function createNVAInVNetWithPrivateNicPublicNicAdditionalPrivateNic() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkVirtualAppliances.createOrUpdate("rg1", "nva", {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1":
          {},
      },
    },
    location: "West US",
    bootStrapConfigurationBlobs: [
      "https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig",
    ],
    cloudInitConfigurationBlobs: [
      "https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig",
    ],
    nvaInterfaceConfigurations: [
      {
        name: "dataInterface",
        type: ["PrivateNic"],
        subnet: {
          id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1",
        },
      },
      {
        name: "managementInterface",
        type: ["PublicNic"],
        subnet: {
          id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet2",
        },
      },
      {
        name: "myAdditionalInterface",
        type: ["AdditionalPrivateNic"],
        subnet: {
          id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet3",
        },
      },
    ],
    nvaSku: { bundledScaleUnit: "1", marketPlaceVersion: "latest", vendor: "Cisco SDWAN" },
    virtualApplianceAsn: 10000,
    tags: { key1: "value1" },
  });
  console.log(result);
}
