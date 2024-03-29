const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a workspace SQL active directory admin
 *
 * @summary Creates or updates a workspace SQL active directory admin
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateWorkspaceAadAdmin.json
 */
async function createOrUpdateWorkspaceActiveDirectoryAdmin() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "resourceGroup1";
  const workspaceName = "workspace1";
  const aadAdminInfo = {
    administratorType: "ActiveDirectory",
    login: "bob@contoso.com",
    sid: "c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c",
    tenantId: "c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.workspaceSqlAadAdmins.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    aadAdminInfo
  );
  console.log(result);
}
