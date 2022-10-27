const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a script package available to run on a private cloud
 *
 * @summary Get a script package available to run on a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/ScriptPackages_Get.json
 */
async function scriptPackagesGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "{privateCloudName}";
  const scriptPackageName = "{scriptPackageName}";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.scriptPackages.get(
    resourceGroupName,
    privateCloudName,
    scriptPackageName
  );
  console.log(result);
}

scriptPackagesGet().catch(console.error);
