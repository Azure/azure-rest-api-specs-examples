const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a WorkloadNetworkPortMirroring
 *
 * @summary Delete a WorkloadNetworkPortMirroring
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_DeletePortMirroring.json
 */
async function workloadNetworksDeletePortMirroring() {
  const subscriptionId =
    process.env["AVS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["AVS_RESOURCE_GROUP"] || "group1";
  const portMirroringId = "portMirroring1";
  const privateCloudName = "cloud1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.workloadNetworks.beginDeletePortMirroringAndWait(
    resourceGroupName,
    portMirroringId,
    privateCloudName,
  );
  console.log(result);
}
