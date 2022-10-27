const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update dhcp by id in a private cloud workload network.
 *
 * @summary Create or update dhcp by id in a private cloud workload network.
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_UpdateDhcpConfigurations.json
 */
async function workloadNetworksUpdateDhcp() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const dhcpId = "dhcp1";
  const workloadNetworkDhcp = {
    properties: {
      dhcpType: "SERVER",
      leaseTime: 86400,
      revision: 1,
      serverAddress: "40.1.5.1/24",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.workloadNetworks.beginUpdateDhcpAndWait(
    resourceGroupName,
    privateCloudName,
    dhcpId,
    workloadNetworkDhcp
  );
  console.log(result);
}

workloadNetworksUpdateDhcp().catch(console.error);
