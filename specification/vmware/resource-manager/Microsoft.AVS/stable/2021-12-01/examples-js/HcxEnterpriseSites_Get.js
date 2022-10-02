const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get an HCX Enterprise Site by name in a private cloud
 *
 * @summary Get an HCX Enterprise Site by name in a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/HcxEnterpriseSites_Get.json
 */
async function hcxEnterpriseSitesGet() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const hcxEnterpriseSiteName = "site1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.hcxEnterpriseSites.get(
    resourceGroupName,
    privateCloudName,
    hcxEnterpriseSiteName
  );
  console.log(result);
}

hcxEnterpriseSitesGet().catch(console.error);
