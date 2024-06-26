const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new workspace.
 *
 * @summary Creates a new workspace.
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/WorkspaceManagedDiskEncryptionUpdate.json
 */
async function updateAWorkspaceWithCustomerManagedKeyCmkEncryptionForManagedDisks() {
  const subscriptionId = process.env["DATABRICKS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["DATABRICKS_RESOURCE_GROUP"] || "rg";
  const workspaceName = "myWorkspace";
  const parameters = {
    encryption: {
      entities: {
        managedDisk: {
          keySource: "Microsoft.Keyvault",
          keyVaultProperties: {
            keyName: "test-cmk-key",
            keyVaultUri: "https://test-vault-name.vault.azure.net/",
            keyVersion: "00000000000000000000000000000000",
          },
          rotationToLatestKeyVersionEnabled: true,
        },
      },
    },
    location: "westus",
    managedResourceGroupId: "/subscriptions/subid/resourceGroups/myManagedRG",
    tags: { mytag1: "myvalue1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.workspaces.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    parameters
  );
  console.log(result);
}
