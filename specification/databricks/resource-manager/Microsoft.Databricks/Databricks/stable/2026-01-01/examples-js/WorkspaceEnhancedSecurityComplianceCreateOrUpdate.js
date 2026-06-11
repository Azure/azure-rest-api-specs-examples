const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new workspace.
 *
 * @summary creates a new workspace.
 * x-ms-original-file: 2026-01-01/WorkspaceEnhancedSecurityComplianceCreateOrUpdate.json
 */
async function createOrUpdateAWorkspaceWithEnhancedSecurityComplianceAddOn() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "11111111-1111-1111-1111-111111111111";
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.workspaces.createOrUpdate("rg", "myWorkspace", {
    location: "eastus2",
    computeMode: "Hybrid",
    enhancedSecurityCompliance: {
      automaticClusterUpdate: { value: "Enabled" },
      complianceSecurityProfile: { complianceStandards: ["PCI_DSS", "HIPAA"], value: "Enabled" },
      enhancedSecurityMonitoring: { value: "Enabled" },
    },
    managedResourceGroupId:
      "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myManagedRG",
    sku: { name: "premium" },
  });
  console.log(result);
}
