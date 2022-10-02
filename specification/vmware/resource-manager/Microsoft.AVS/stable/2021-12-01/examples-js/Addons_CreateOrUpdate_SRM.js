const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a addon in a private cloud
 *
 * @summary Create or update a addon in a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Addons_CreateOrUpdate_SRM.json
 */
async function addonsCreateOrUpdateSrm() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const addonName = "srm";
  const addon = {
    properties: {
      addonType: "SRM",
      licenseKey: "41915178-A8FF-4A4D-B683-6D735AF5E3F5",
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

addonsCreateOrUpdateSrm().catch(console.error);
