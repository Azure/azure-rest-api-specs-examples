const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a addon in a private cloud
 *
 * @summary Delete a addon in a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/Addons_Delete.json
 */
async function addonsDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const addonName = "srm";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.addons.beginDeleteAndWait(
    resourceGroupName,
    privateCloudName,
    addonName
  );
  console.log(result);
}

addonsDelete().catch(console.error);
