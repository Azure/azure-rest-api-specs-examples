const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a addon in a private cloud
 *
 * @summary Create or update a addon in a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/Addons_CreateOrUpdate_VR.json
 */
async function addonsCreateOrUpdateVr() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const addonName = "vr";
  const addon = { properties: { addonType: "VR", vrsCount: 1 } };
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

addonsCreateOrUpdateVr().catch(console.error);
