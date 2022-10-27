const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List script cmdlet resources available for a private cloud to create a script execution resource on a private cloud
 *
 * @summary List script cmdlet resources available for a private cloud to create a script execution resource on a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/ScriptCmdlets_List.json
 */
async function scriptCmdletsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "{privateCloudName}";
  const scriptPackageName = "{scriptPackageName}";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.scriptCmdlets.list(
    resourceGroupName,
    privateCloudName,
    scriptPackageName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

scriptCmdletsList().catch(console.error);
