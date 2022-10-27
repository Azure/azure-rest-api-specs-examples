const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List script packages available to run on the private cloud
 *
 * @summary List script packages available to run on the private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/ScriptPackages_List.json
 */
async function scriptPackagesList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "{privateCloudName}";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.scriptPackages.list(resourceGroupName, privateCloudName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

scriptPackagesList().catch(console.error);
