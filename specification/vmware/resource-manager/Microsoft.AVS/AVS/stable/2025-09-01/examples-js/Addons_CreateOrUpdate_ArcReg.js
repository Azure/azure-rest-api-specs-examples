const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a Addon
 *
 * @summary create a Addon
 * x-ms-original-file: 2025-09-01/Addons_CreateOrUpdate_ArcReg.json
 */
async function addonsCreateOrUpdateArcReg() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.addons.createOrUpdate("group1", "cloud1", "arc", {
    properties: {
      addonType: "Arc",
      vCenter:
        "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg_test/providers/Microsoft.ConnectedVMwarevSphere/VCenters/test-vcenter",
    },
  });
  console.log(result);
}
