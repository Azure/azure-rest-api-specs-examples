using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ManagedPrivateEndpoints_Get.json
// this example is just showing the usage of "ManagedPrivateEndpoints_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryManagedVirtualNetworkResource created on azure
// for more information of creating DataFactoryManagedVirtualNetworkResource, please refer to the document of DataFactoryManagedVirtualNetworkResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string managedVirtualNetworkName = "exampleManagedVirtualNetworkName";
ResourceIdentifier dataFactoryManagedVirtualNetworkResourceId = DataFactoryManagedVirtualNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, managedVirtualNetworkName);
DataFactoryManagedVirtualNetworkResource dataFactoryManagedVirtualNetwork = client.GetDataFactoryManagedVirtualNetworkResource(dataFactoryManagedVirtualNetworkResourceId);

// get the collection of this DataFactoryPrivateEndpointResource
DataFactoryPrivateEndpointCollection collection = dataFactoryManagedVirtualNetwork.GetDataFactoryPrivateEndpoints();

// invoke the operation
string managedPrivateEndpointName = "exampleManagedPrivateEndpointName";
NullableResponse<DataFactoryPrivateEndpointResource> response = await collection.GetIfExistsAsync(managedPrivateEndpointName);
DataFactoryPrivateEndpointResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DataFactoryPrivateEndpointData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
