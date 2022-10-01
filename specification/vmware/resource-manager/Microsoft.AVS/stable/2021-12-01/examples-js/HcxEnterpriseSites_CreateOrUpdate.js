const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an HCX Enterprise Site in a private cloud
 *
 * @summary Create or update an HCX Enterprise Site in a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/HcxEnterpriseSites_CreateOrUpdate.json
 */
async function hcxEnterpriseSitesCreateOrUpdate() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const hcxEnterpriseSiteName = "site1";
  const hcxEnterpriseSite = {};
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.hcxEnterpriseSites.createOrUpdate(
    resourceGroupName,
    privateCloudName,
    hcxEnterpriseSiteName,
    hcxEnterpriseSite
  );
  console.log(result);
}

hcxEnterpriseSitesCreateOrUpdate().catch(console.error);
