const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Settings about where we should store your security data and logs
 *
 * @summary Settings about where we should store your security data and logs
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/WorkspaceSettings/UpdateWorkspaceSetting_example.json
 */
async function updateAWorkspaceSettingDataForSubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const workspaceSettingName = "default";
  const workspaceSetting = {
    name: "default",
    type: "Microsoft.Security/workspaceSettings",
    id: "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/workspaceSettings/default",
    workspaceId:
      "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.workspaceSettings.update(workspaceSettingName, workspaceSetting);
  console.log(result);
}

updateAWorkspaceSettingDataForSubscription().catch(console.error);
