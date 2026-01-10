const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a PrivateCloud
 *
 * @summary create a PrivateCloud
 * x-ms-original-file: 2025-09-01/PrivateClouds_CreateOrUpdate_FleetNative.json
 */
async function privateCloudsCreateOrUpdateFleetNative() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.privateClouds.createOrUpdate("group1", "cloud1", {
    location: "eastus2",
    sku: { name: "AV64" },
    properties: {
      networkBlock: "192.168.48.0/22",
      virtualNetworkId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualNetworks/vnet",
      dnsZoneType: "Private",
      managementCluster: { clusterSize: 4 },
      vcfLicense: {
        kind: "vcf5",
        licenseKey: "12345-12345-12345-12345-12345",
        endDate: new Date("2025-12-31T23:59:59Z"),
        cores: 16,
        broadcomSiteId: "123456",
        broadcomContractNumber: "123456",
      },
    },
    tags: {},
  });
  console.log(result);
}
