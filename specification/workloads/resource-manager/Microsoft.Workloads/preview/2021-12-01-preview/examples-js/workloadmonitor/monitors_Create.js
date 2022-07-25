const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a SAP monitor for the specified subscription, resource group, and resource name.
 *
 * @summary Creates a SAP monitor for the specified subscription, resource group, and resource name.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/monitors_Create.json
 */
async function createASapMonitor() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "mySapMonitor";
  const monitorParameter = {
    appLocation: "westus",
    location: "westus",
    logAnalyticsWorkspaceArmId:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.operationalinsights/workspaces/myWorkspace",
    managedResourceGroupConfiguration: { name: "myManagedRg" },
    monitorSubnet:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet",
    routingPreference: "RouteAll",
    tags: { key: "value" },
  };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.monitors.beginCreateAndWait(
    resourceGroupName,
    monitorName,
    monitorParameter
  );
  console.log(result);
}

createASapMonitor().catch(console.error);
