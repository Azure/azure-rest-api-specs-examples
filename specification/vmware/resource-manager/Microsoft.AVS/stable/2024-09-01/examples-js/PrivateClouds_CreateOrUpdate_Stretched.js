const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a PrivateCloud
 *
 * @summary create a PrivateCloud
 * x-ms-original-file: 2024-09-01/PrivateClouds_CreateOrUpdate_Stretched.json
 */
async function privateCloudsCreateOrUpdateStretched() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.privateClouds.createOrUpdate("group1", "cloud1", {
    location: "eastus2",
    sku: { name: "AV36" },
    properties: {
      availability: { strategy: "DualZone", zone: 1, secondaryZone: 2 },
      networkBlock: "192.168.48.0/22",
      managementCluster: { clusterSize: 4 },
    },
    tags: {},
  });
  console.log(result);
}
