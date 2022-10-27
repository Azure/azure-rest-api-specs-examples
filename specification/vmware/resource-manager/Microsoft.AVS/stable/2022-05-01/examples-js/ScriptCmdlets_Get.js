const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return information about a script cmdlet resource in a specific package on a private cloud
 *
 * @summary Return information about a script cmdlet resource in a specific package on a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/ScriptCmdlets_Get.json
 */
async function scriptCmdletsGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "{privateCloudName}";
  const scriptPackageName = "{scriptPackageName}";
  const scriptCmdletName = "New-ExternalSsoDomain";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.scriptCmdlets.get(
    resourceGroupName,
    privateCloudName,
    scriptPackageName,
    scriptCmdletName
  );
  console.log(result);
}

scriptCmdletsGet().catch(console.error);
