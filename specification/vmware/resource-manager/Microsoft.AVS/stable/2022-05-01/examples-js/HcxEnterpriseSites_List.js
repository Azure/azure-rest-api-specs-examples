const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List HCX Enterprise Sites in a private cloud
 *
 * @summary List HCX Enterprise Sites in a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/HcxEnterpriseSites_List.json
 */
async function hcxEnterpriseSitesList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.hcxEnterpriseSites.list(resourceGroupName, privateCloudName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

hcxEnterpriseSitesList().catch(console.error);
