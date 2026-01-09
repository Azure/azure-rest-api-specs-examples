const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a PrivateCloud
 *
 * @summary get a PrivateCloud
 * x-ms-original-file: 2025-09-01/PrivateClouds_Get_Stretched.json
 */
async function privateCloudsGetStretched() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.privateClouds.get("group1", "cloud1");
  console.log(result);
}
