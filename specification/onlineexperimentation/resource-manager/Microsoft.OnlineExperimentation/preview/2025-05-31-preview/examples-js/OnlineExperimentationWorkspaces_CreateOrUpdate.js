const { OnlineExperimentationClient } = require("@azure/arm-onlineexperimentation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create an online experimentation workspace, or update an existing workspace.
 *
 * @summary create an online experimentation workspace, or update an existing workspace.
 * x-ms-original-file: 2025-05-31-preview/OnlineExperimentationWorkspaces_CreateOrUpdate.json
 */
async function createOrUpdateAnOnlineExperimentationWorkspaceWithFreeSku() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "fa5fc227-a624-475e-b696-cdd604c735bc";
  const client = new OnlineExperimentationClient(credential, subscriptionId);
  const result = await client.onlineExperimentationWorkspaces.createOrUpdate(
    "res9871",
    "expworkspace7",
    {
      location: "eastus2",
      tags: { newKey: "newVal" },
      properties: {
        logAnalyticsWorkspaceResourceId:
          "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.OperationalInsights/workspaces/log9871",
        logsExporterStorageAccountResourceId:
          "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.Storage/storageAccounts/sto9871",
        appConfigurationResourceId:
          "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/res9871/providers/Microsoft.AppConfiguration/configurationStores/appconfig9871",
      },
      identity: {
        type: "UserAssigned",
        userAssignedIdentities: {
          "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1":
            {},
          "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2":
            {},
        },
      },
      sku: { name: "F0" },
    },
  );
  console.log(result);
}
