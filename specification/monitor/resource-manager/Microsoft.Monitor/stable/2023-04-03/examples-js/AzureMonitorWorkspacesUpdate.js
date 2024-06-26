const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates part of an Azure Monitor Workspace
 *
 * @summary Updates part of an Azure Monitor Workspace
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Monitor/stable/2023-04-03/examples/AzureMonitorWorkspacesUpdate.json
 */
async function updateAnAzureMonitorWorkspace() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "703362b3-f278-4e4b-9179-c76eaf41ffc2";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "myResourceGroup";
  const azureMonitorWorkspaceName = "myAzureMonitorWorkspace";
  const azureMonitorWorkspaceProperties = { tags: { tag1: "A", tag2: "B", tag3: "C" } };
  const options = {
    azureMonitorWorkspaceProperties,
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.azureMonitorWorkspaces.update(
    resourceGroupName,
    azureMonitorWorkspaceName,
    options,
  );
  console.log(result);
}
