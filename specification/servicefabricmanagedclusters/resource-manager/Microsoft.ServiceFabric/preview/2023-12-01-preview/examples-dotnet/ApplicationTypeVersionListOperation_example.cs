using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceFabricManagedClusters;

// Generated from example definition: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2023-12-01-preview/examples/ApplicationTypeVersionListOperation_example.json
// this example is just showing the usage of "ApplicationTypeVersions_ListByApplicationTypes" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceFabricManagedApplicationTypeResource created on azure
// for more information of creating ServiceFabricManagedApplicationTypeResource, please refer to the document of ServiceFabricManagedApplicationTypeResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string clusterName = "myCluster";
string applicationTypeName = "myAppType";
ResourceIdentifier serviceFabricManagedApplicationTypeResourceId = ServiceFabricManagedApplicationTypeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, applicationTypeName);
ServiceFabricManagedApplicationTypeResource serviceFabricManagedApplicationType = client.GetServiceFabricManagedApplicationTypeResource(serviceFabricManagedApplicationTypeResourceId);

// get the collection of this ServiceFabricManagedApplicationTypeVersionResource
ServiceFabricManagedApplicationTypeVersionCollection collection = serviceFabricManagedApplicationType.GetServiceFabricManagedApplicationTypeVersions();

// invoke the operation and iterate over the result
await foreach (ServiceFabricManagedApplicationTypeVersionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ServiceFabricManagedApplicationTypeVersionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
