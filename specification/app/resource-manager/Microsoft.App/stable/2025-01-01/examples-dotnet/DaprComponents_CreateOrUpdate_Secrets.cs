using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/DaprComponents_CreateOrUpdate_Secrets.json
// this example is just showing the usage of "DaprComponents_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppManagedEnvironmentDaprComponentResource created on azure
// for more information of creating ContainerAppManagedEnvironmentDaprComponentResource, please refer to the document of ContainerAppManagedEnvironmentDaprComponentResource
string subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
string resourceGroupName = "examplerg";
string environmentName = "myenvironment";
string componentName = "reddog";
ResourceIdentifier containerAppManagedEnvironmentDaprComponentResourceId = ContainerAppManagedEnvironmentDaprComponentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, environmentName, componentName);
ContainerAppManagedEnvironmentDaprComponentResource containerAppManagedEnvironmentDaprComponent = client.GetContainerAppManagedEnvironmentDaprComponentResource(containerAppManagedEnvironmentDaprComponentResourceId);

// invoke the operation
ContainerAppDaprComponentData data = new ContainerAppDaprComponentData
{
    ComponentType = "state.azure.cosmosdb",
    Version = "v1",
    IgnoreErrors = false,
    InitTimeout = "50s",
    Secrets = {new ContainerAppWritableSecret
    {
    Name = "masterkey",
    Value = "keyvalue",
    }},
    Metadata = {new ContainerAppDaprMetadata
    {
    Name = "url",
    Value = "<COSMOS-URL>",
    }, new ContainerAppDaprMetadata
    {
    Name = "database",
    Value = "itemsDB",
    }, new ContainerAppDaprMetadata
    {
    Name = "collection",
    Value = "items",
    }, new ContainerAppDaprMetadata
    {
    Name = "masterkey",
    SecretRef = "masterkey",
    }},
    Scopes = { "container-app-1", "container-app-2" },
};
ArmOperation<ContainerAppManagedEnvironmentDaprComponentResource> lro = await containerAppManagedEnvironmentDaprComponent.UpdateAsync(WaitUntil.Completed, data);
ContainerAppManagedEnvironmentDaprComponentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerAppDaprComponentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
