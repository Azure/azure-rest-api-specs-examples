const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a workspace
 *
 * @summary Create or update a workspace
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Monitor/preview/2021-06-03-preview/examples/AzureMonitorWorkspacesCreate.json
 */
async function createOrUpdateWorkspace() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "703362b3-f278-4e4b-9179-c76eaf41ffc2";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "myResourceGroup";
  const azureMonitorWorkspaceName = "myAzureMonitorWorkspace";
  const azureMonitorWorkspaceProperties = {
    location: "eastus",
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.azureMonitorWorkspaces.create(
    resourceGroupName,
    azureMonitorWorkspaceName,
    azureMonitorWorkspaceProperties
  );
  console.log(result);
}
