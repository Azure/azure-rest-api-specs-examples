const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a addon in a private cloud
 *
 * @summary Create or update a addon in a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/Addons_CreateOrUpdate_ArcReg.json
 */
async function addonsCreateOrUpdateArc() {
  const subscriptionId =
    process.env["AVS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["AVS_RESOURCE_GROUP"] || "group1";
  const privateCloudName = "cloud1";
  const addonName = "arc";
  const addon = {
    properties: {
      addonType: "Arc",
      vCenter:
        "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg_test/providers/Microsoft.ConnectedVMwarevSphere/VCenters/test-vcenter",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.addons.beginCreateOrUpdateAndWait(
    resourceGroupName,
    privateCloudName,
    addonName,
    addon
  );
  console.log(result);
}
