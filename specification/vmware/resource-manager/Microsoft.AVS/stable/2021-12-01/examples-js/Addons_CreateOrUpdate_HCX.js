const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a addon in a private cloud
 *
 * @summary Create or update a addon in a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Addons_CreateOrUpdate_HCX.json
 */
async function addonsCreateOrUpdateHcx() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const addonName = "hcx";
  const addon = {
    properties: {
      addonType: "HCX",
      offer: "VMware MaaS Cloud Provider (Enterprise)",
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

addonsCreateOrUpdateHcx().catch(console.error);
