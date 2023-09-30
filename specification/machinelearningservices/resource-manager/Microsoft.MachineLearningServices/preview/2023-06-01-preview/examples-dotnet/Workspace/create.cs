using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Workspace/create.json
// this example is just showing the usage of "Workspaces_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "workspace-1234";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this MachineLearningWorkspaceResource
MachineLearningWorkspaceCollection collection = resourceGroupResource.GetMachineLearningWorkspaces();

// invoke the operation
string workspaceName = "testworkspace";
MachineLearningWorkspaceData data = new MachineLearningWorkspaceData(new AzureLocation("eastus2euap"))
{
    Identity = new ManagedServiceIdentity("SystemAssigned,UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai")] = new UserAssignedIdentity(),
        },
    },
    ApplicationInsights = "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/microsoft.insights/components/testinsights",
    ContainerRegistry = "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ContainerRegistry/registries/testRegistry",
    Description = "test description",
    Encryption = new MachineLearningEncryptionSetting(new MachineLearningEncryptionKeyVaultProperties("https://testkv.vault.azure.net/keys/testkey/aabbccddee112233445566778899aabb", new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv"))
    {
        IdentityClientId = "",
    }, MachineLearningEncryptionStatus.Enabled)
    {
        UserAssignedIdentity = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai"),
    },
    FriendlyName = "HelloName",
    IsHbiWorkspace = false,
    KeyVault = "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv",
    SharedPrivateLinkResources =
    {
    new MachineLearningSharedPrivateLinkResource()
    {
    Name = "testdbresource",
    GroupId = "Sql",
    PrivateLinkResourceId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.DocumentDB/databaseAccounts/testdbresource/privateLinkResources/Sql"),
    RequestMessage = "Please approve",
    Status = MachineLearningPrivateEndpointServiceConnectionStatus.Approved,
    }
    },
    StorageAccount = "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.Storage/storageAccounts/testStorageAccount",
};
ArmOperation<MachineLearningWorkspaceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, workspaceName, data);
MachineLearningWorkspaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningWorkspaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
