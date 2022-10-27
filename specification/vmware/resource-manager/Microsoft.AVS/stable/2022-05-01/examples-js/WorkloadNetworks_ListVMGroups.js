const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of vm groups in a private cloud workload network.
 *
 * @summary List of vm groups in a private cloud workload network.
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_ListVMGroups.json
 */
async function workloadNetworksListVMGroups() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workloadNetworks.listVMGroups(
    resourceGroupName,
    privateCloudName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

workloadNetworksListVMGroups().catch(console.error);
