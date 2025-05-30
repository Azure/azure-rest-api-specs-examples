using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Synapse.Models;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateWorkspace.json
// this example is just showing the usage of "Workspaces_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "resourceGroup1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this SynapseWorkspaceResource
SynapseWorkspaceCollection collection = resourceGroupResource.GetSynapseWorkspaces();

// invoke the operation
string workspaceName = "workspace1";
SynapseWorkspaceData data = new SynapseWorkspaceData(new AzureLocation("East US"))
{
    Identity = new ManagedServiceIdentity("SystemAssigned,UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1")] = new UserAssignedIdentity()
        },
    },
    DefaultDataLakeStorage = new SynapseDataLakeStorageAccountDetails
    {
        AccountUri = new Uri("https://accountname.dfs.core.windows.net"),
        Filesystem = "default",
    },
    SqlAdministratorLoginPassword = "password",
    ManagedResourceGroupName = "workspaceManagedResourceGroupUnique",
    SqlAdministratorLogin = "login",
    ManagedVirtualNetwork = "default",
    Encryption = new SynapseEncryptionDetails
    {
        Cmk = new WorkspaceCustomerManagedKeyDetails
        {
            Key = new SynapseWorkspaceKeyDetails
            {
                Name = "default",
                KeyVaultUri = new Uri("https://vault.azure.net/keys/key1"),
            },
            KekIdentity = new KekIdentityProperties
            {
                UserAssignedIdentityId = new ResourceIdentifier("/subscriptions/b64d7b94-73e7-4d36-94b2-7764ea3fd74a/resourcegroups/SynapseCI/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
                UseSystemAssignedIdentity = BinaryData.FromObjectAsJson("false"),
            },
        },
    },
    ManagedVirtualNetworkSettings = new SynapseManagedVirtualNetworkSettings
    {
        PreventDataExfiltration = false,
        EnableLinkedAccessCheckOnTargetResource = false,
        AllowedAadTenantIdsForLinking = { "740239CE-A25B-485B-86A0-262F29F6EBDB" },
    },
    WorkspaceRepositoryConfiguration = new SynapseWorkspaceRepositoryConfiguration
    {
        WorkspaceRepositoryConfigurationType = "FactoryGitHubConfiguration",
        HostName = "",
        AccountName = "mygithubaccount",
        ProjectName = "myproject",
        RepositoryName = "myrepository",
        CollaborationBranch = "master",
        RootFolder = "/",
    },
    PurviewResourceId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1"),
    PublicNetworkAccess = WorkspacePublicNetworkAccess.Enabled,
    InitialWorkspaceAdminObjectId = Guid.Parse("6c20646f-8050-49ec-b3b1-80a0e58e454d"),
    Tags =
    {
    ["key"] = "value"
    },
};
ArmOperation<SynapseWorkspaceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, workspaceName, data);
SynapseWorkspaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseWorkspaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
