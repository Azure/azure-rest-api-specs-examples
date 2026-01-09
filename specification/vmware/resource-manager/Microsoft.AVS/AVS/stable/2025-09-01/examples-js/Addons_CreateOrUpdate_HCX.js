const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a Addon
 *
 * @summary create a Addon
 * x-ms-original-file: 2025-09-01/Addons_CreateOrUpdate_HCX.json
 */
async function addonsCreateOrUpdateHCX() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.addons.createOrUpdate("group1", "cloud1", "hcx", {
    properties: { addonType: "HCX", offer: "VMware MaaS Cloud Provider (Enterprise)" },
  });
  console.log(result);
}
