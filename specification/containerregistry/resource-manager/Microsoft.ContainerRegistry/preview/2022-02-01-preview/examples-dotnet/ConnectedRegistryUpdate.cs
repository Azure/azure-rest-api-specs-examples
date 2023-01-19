using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerRegistry;
using Azure.ResourceManager.ContainerRegistry.Models;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/ConnectedRegistryUpdate.json
// this example is just showing the usage of "ConnectedRegistries_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConnectedRegistryResource created on azure
// for more information of creating ConnectedRegistryResource, please refer to the document of ConnectedRegistryResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
string connectedRegistryName = "myScopeMap";
ResourceIdentifier connectedRegistryResourceId = ConnectedRegistryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, connectedRegistryName);
ConnectedRegistryResource connectedRegistry = client.GetConnectedRegistryResource(connectedRegistryResourceId);

// invoke the operation
ConnectedRegistryPatch patch = new ConnectedRegistryPatch()
{
    SyncProperties = new ConnectedRegistrySyncUpdateProperties()
    {
        Schedule = "0 0 */10 * *",
        SyncWindow = XmlConvert.ToTimeSpan("P2D"),
        MessageTtl = XmlConvert.ToTimeSpan("P30D"),
    },
    Logging = new ConnectedRegistryLogging()
    {
        LogLevel = ConnectedRegistryLogLevel.Debug,
        AuditLogStatus = ConnectedRegistryAuditLogStatus.Enabled,
    },
    ClientTokenIds =
    {
    new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token"),new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client2Token")
    },
    NotificationsList =
    {
    "hello-world:*:*","sample/repo/*:1.0:*"
    },
};
ArmOperation<ConnectedRegistryResource> lro = await connectedRegistry.UpdateAsync(WaitUntil.Completed, patch);
ConnectedRegistryResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ConnectedRegistryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
