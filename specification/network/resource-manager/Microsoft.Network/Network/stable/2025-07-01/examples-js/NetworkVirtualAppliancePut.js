const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the specified Network Virtual Appliance.
 *
 * @summary creates or updates the specified Network Virtual Appliance.
 * x-ms-original-file: 2025-07-01/NetworkVirtualAppliancePut.json
 */
async function createNetworkVirtualAppliance() {
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
    additionalNics: [{ name: "exrsdwan", hasPublicIp: true }],
    bootStrapConfigurationBlobs: [
      "https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig",
    ],
    cloudInitConfigurationBlobs: [
      "https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig",
    ],
    internetIngressPublicIps: [
      {
        id: "/subscriptions/{{subscriptionId}}/resourceGroups/{{rg}}/providers/Microsoft.Network/publicIPAddresses/slbip",
      },
    ],
    networkProfile: {
      networkInterfaceConfigurations: [
        {
          nicType: "PublicNic",
          properties: {
            ipConfigurations: [
              { name: "publicnicipconfig", properties: { primary: true } },
              { name: "publicnicipconfig-2", properties: { primary: false } },
            ],
          },
        },
        {
          nicType: "PrivateNic",
          properties: {
            ipConfigurations: [
              { name: "privatenicipconfig", properties: { primary: true } },
              { name: "privatenicipconfig-2", properties: { primary: false } },
            ],
          },
        },
      ],
    },
    nvaSku: { bundledScaleUnit: "1", marketPlaceVersion: "12.1", vendor: "Cisco SDWAN" },
    virtualApplianceAsn: 10000,
    virtualHub: {
      id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1",
    },
    tags: { key1: "value1" },
  });
  console.log(result);
}
