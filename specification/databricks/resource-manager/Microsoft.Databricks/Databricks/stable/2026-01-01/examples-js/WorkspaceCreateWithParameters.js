const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new workspace.
 *
 * @summary creates a new workspace.
 * x-ms-original-file: 2026-01-01/WorkspaceCreateWithParameters.json
 */
async function createOrUpdateWorkspaceWithCustomParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "11111111-1111-1111-1111-111111111111";
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.workspaces.createOrUpdate("rg", "myWorkspace", {
    location: "westus",
    accessConnector: {
      id: "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/adbrg/providers/Microsoft.Databricks/accessConnectors/myAccessConnector",
      identityType: "UserAssigned",
      userAssignedIdentityId:
        "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity",
    },
    computeMode: "Hybrid",
    defaultCatalog: { initialName: "", initialType: "HiveMetastore" },
    defaultStorageFirewall: "Enabled",
    managedResourceGroupId:
      "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myManagedRG",
    parameters: {
      customPrivateSubnetName: { value: "myPrivateSubnet" },
      customPublicSubnetName: { value: "myPublicSubnet" },
      customVirtualNetworkId: {
        value:
          "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/myNetwork",
      },
    },
    sku: { name: "premium" },
  });
  console.log(result);
}
