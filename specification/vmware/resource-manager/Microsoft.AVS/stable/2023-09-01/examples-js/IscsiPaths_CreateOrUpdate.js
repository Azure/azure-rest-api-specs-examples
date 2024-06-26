const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a IscsiPath
 *
 * @summary Create a IscsiPath
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/IscsiPaths_CreateOrUpdate.json
 */
async function iscsiPathsCreateOrUpdate() {
  const subscriptionId =
    process.env["AVS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["AVS_RESOURCE_GROUP"] || "group1";
  const privateCloudName = "cloud1";
  const resource = { networkBlock: "192.168.0.0/24" };
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.iscsiPaths.beginCreateOrUpdateAndWait(
    resourceGroupName,
    privateCloudName,
    resource,
  );
  console.log(result);
}
