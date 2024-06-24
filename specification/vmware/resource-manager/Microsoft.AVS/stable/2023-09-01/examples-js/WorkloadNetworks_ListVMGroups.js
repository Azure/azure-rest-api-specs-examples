const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List WorkloadNetworkVMGroup resources by WorkloadNetwork
 *
 * @summary List WorkloadNetworkVMGroup resources by WorkloadNetwork
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_ListVMGroups.json
 */
async function workloadNetworksListVMGroups() {
  const subscriptionId =
    process.env["AVS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["AVS_RESOURCE_GROUP"] || "group1";
  const privateCloudName = "cloud1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workloadNetworks.listVMGroups(
    resourceGroupName,
    privateCloudName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
