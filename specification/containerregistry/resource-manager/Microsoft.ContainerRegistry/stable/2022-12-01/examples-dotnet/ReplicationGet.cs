using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerRegistry;
using Azure.ResourceManager.ContainerRegistry.Models;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2022-12-01/examples/ReplicationGet.json
// this example is just showing the usage of "Replications_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerRegistryReplicationResource created on azure
// for more information of creating ContainerRegistryReplicationResource, please refer to the document of ContainerRegistryReplicationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
string replicationName = "myReplication";
ResourceIdentifier containerRegistryReplicationResourceId = ContainerRegistryReplicationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, replicationName);
ContainerRegistryReplicationResource containerRegistryReplication = client.GetContainerRegistryReplicationResource(containerRegistryReplicationResourceId);

// invoke the operation
ContainerRegistryReplicationResource result = await containerRegistryReplication.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryReplicationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
