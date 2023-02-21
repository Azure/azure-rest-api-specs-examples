using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateWorkspace.json
// this example is just showing the usage of "Workspaces_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseWorkspaceResource created on azure
// for more information of creating SynapseWorkspaceResource, please refer to the document of SynapseWorkspaceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "resourceGroup1";
string workspaceName = "workspace1";
ResourceIdentifier synapseWorkspaceResourceId = SynapseWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
SynapseWorkspaceResource synapseWorkspace = client.GetSynapseWorkspaceResource(synapseWorkspaceResourceId);

// invoke the operation
SynapseWorkspacePatch patch = new SynapseWorkspacePatch()
{
    Tags =
    {
    ["key"] = "value",
    },
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    SqlAdministratorLoginPassword = "password",
    ManagedVirtualNetworkSettings = new SynapseManagedVirtualNetworkSettings()
    {
        PreventDataExfiltration = false,
        EnableLinkedAccessCheckOnTargetResource = false,
        AllowedAadTenantIdsForLinking =
        {
        "740239CE-A25B-485B-86A0-262F29F6EBDB"
        },
    },
    WorkspaceRepositoryConfiguration = new SynapseWorkspaceRepositoryConfiguration()
    {
        WorkspaceRepositoryConfigurationType = "FactoryGitHubConfiguration",
        HostName = "",
        AccountName = "adifferentacount",
        ProjectName = "myproject",
        RepositoryName = "myrepository",
        CollaborationBranch = "master",
        RootFolder = "/",
    },
    PurviewResourceId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1"),
    Encryption = new SynapseEncryptionDetails()
    {
        Cmk = new WorkspaceCustomerManagedKeyDetails()
        {
            Key = new SynapseWorkspaceKeyDetails()
            {
                Name = "default",
                KeyVaultUri = new Uri("https://vault.azure.net/keys/key1"),
            },
        },
    },
    PublicNetworkAccess = WorkspacePublicNetworkAccess.Enabled,
};
ArmOperation<SynapseWorkspaceResource> lro = await synapseWorkspace.UpdateAsync(WaitUntil.Completed, patch);
SynapseWorkspaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseWorkspaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
