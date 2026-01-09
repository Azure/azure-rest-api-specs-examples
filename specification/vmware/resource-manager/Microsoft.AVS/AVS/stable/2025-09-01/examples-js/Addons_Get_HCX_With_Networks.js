const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a Addon
 *
 * @summary get a Addon
 * x-ms-original-file: 2025-09-01/Addons_Get_HCX_With_Networks.json
 */
async function addonsGetHCXWithNetworks() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.addons.get("group1", "cloud1", "hcx");
  console.log(result);
}
