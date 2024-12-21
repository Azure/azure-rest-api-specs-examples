const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restarts one or more VMs belonging to the specified Network Virtual Appliance.
 *
 * @summary Restarts one or more VMs belonging to the specified Network Virtual Appliance.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkVirtualApplianceEmptyRestart.json
 */
async function restartAllNetworkVirtualApplianceVMSInVMScaleSet() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkVirtualApplianceName = "nva";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkVirtualAppliances.beginRestartAndWait(
    resourceGroupName,
    networkVirtualApplianceName,
  );
  console.log(result);
}
