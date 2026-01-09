const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a PrivateCloud
 *
 * @summary update a PrivateCloud
 * x-ms-original-file: 2025-09-01/PrivateClouds_Update_Stretched.json
 */
async function privateCloudsUpdateStretched() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.privateClouds.update("group1", "cloud1", {
    properties: { managementCluster: { clusterSize: 4 } },
  });
  console.log(result);
}
