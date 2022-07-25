const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a workspace
 *
 * @summary Creates or updates a workspace
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateWorkspace.json
 */
async function createOrUpdateAWorkspace() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "resourceGroup1";
  const workspaceName = "workspace1";
  const workspaceInfo = {
    cspWorkspaceAdminProperties: {
      initialWorkspaceAdminObjectId: "6c20646f-8050-49ec-b3b1-80a0e58e454d",
    },
    defaultDataLakeStorage: {
      accountUrl: "https://accountname.dfs.core.windows.net",
      filesystem: "default",
    },
    encryption: {
      cmk: {
        kekIdentity: {
          useSystemAssignedIdentity: false,
          userAssignedIdentity:
            "/subscriptions/b64d7b94-73e7-4d36-94b2-7764ea3fd74a/resourcegroups/SynapseCI/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1",
        },
        key: {
          name: "default",
          keyVaultUrl: "https://vault.azure.net/keys/key1",
        },
      },
    },
    identity: {
      type: "SystemAssigned,UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000111122223333444444444444/resourcegroups/resourceGroup1/providers/MicrosoftManagedIdentity/userAssignedIdentities/uami1":
          {},
      },
    },
    location: "East US",
    managedResourceGroupName: "workspaceManagedResourceGroupUnique",
    managedVirtualNetwork: "default",
    managedVirtualNetworkSettings: {
      allowedAadTenantIdsForLinking: ["740239CE-A25B-485B-86A0-262F29F6EBDB"],
      linkedAccessCheckOnTargetResource: false,
      preventDataExfiltration: false,
    },
    publicNetworkAccess: "Enabled",
    purviewConfiguration: {
      purviewResourceId:
        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1",
    },
    sqlAdministratorLogin: "login",
    sqlAdministratorLoginPassword: "password",
    tags: { key: "value" },
    workspaceRepositoryConfiguration: {
      type: "FactoryGitHubConfiguration",
      accountName: "mygithubaccount",
      collaborationBranch: "master",
      hostName: "",
      projectName: "myproject",
      repositoryName: "myrepository",
      rootFolder: "/",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.workspaces.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    workspaceInfo
  );
  console.log(result);
}

createOrUpdateAWorkspace().catch(console.error);
