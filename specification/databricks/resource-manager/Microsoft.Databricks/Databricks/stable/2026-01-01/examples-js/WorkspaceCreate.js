const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new workspace.
 *
 * @summary creates a new workspace.
 * x-ms-original-file: 2026-01-01/WorkspaceCreate.json
 */
async function createOrUpdateWorkspace() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "11111111-1111-1111-1111-111111111111";
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.workspaces.createOrUpdate("rg", "myWorkspace", {
    location: "westus",
    accessConnector: {
      id: "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/adbrg/providers/Microsoft.Databricks/accessConnectors/myAccessConnector",
      identityType: "SystemAssigned",
    },
    computeMode: "Hybrid",
    defaultCatalog: { initialName: "", initialType: "UnityCatalog" },
    defaultStorageFirewall: "Enabled",
    managedResourceGroupId:
      "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myManagedRG",
    sku: { name: "premium" },
  });
  console.log(result);
}
