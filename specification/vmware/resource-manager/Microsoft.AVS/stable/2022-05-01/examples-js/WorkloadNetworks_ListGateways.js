const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of gateways in a private cloud workload network.
 *
 * @summary List of gateways in a private cloud workload network.
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_ListGateways.json
 */
async function workloadNetworksListGateways() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workloadNetworks.listGateways(
    resourceGroupName,
    privateCloudName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

workloadNetworksListGateways().catch(console.error);
