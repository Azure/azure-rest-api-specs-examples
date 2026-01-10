const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a Addon
 *
 * @summary create a Addon
 * x-ms-original-file: 2025-09-01/Addons_CreateOrUpdate_VR.json
 */
async function addonsCreateOrUpdateVR() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.addons.createOrUpdate("group1", "cloud1", "vr", {
    properties: { addonType: "VR", vrsCount: 1 },
  });
  console.log(result);
}
