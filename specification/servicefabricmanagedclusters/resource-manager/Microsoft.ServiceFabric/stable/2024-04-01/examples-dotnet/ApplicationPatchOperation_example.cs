using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceFabricManagedClusters.Models;
using Azure.ResourceManager.ServiceFabricManagedClusters;

// Generated from example definition: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/ApplicationPatchOperation_example.json
// this example is just showing the usage of "Applications_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceFabricManagedApplicationResource created on azure
// for more information of creating ServiceFabricManagedApplicationResource, please refer to the document of ServiceFabricManagedApplicationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string clusterName = "myCluster";
string applicationName = "myApp";
ResourceIdentifier serviceFabricManagedApplicationResourceId = ServiceFabricManagedApplicationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, applicationName);
ServiceFabricManagedApplicationResource serviceFabricManagedApplication = client.GetServiceFabricManagedApplicationResource(serviceFabricManagedApplicationResourceId);

// invoke the operation
ServiceFabricManagedApplicationPatch patch = new ServiceFabricManagedApplicationPatch()
{
    Tags =
    {
    ["a"] = "b",
    },
};
ServiceFabricManagedApplicationResource result = await serviceFabricManagedApplication.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceFabricManagedApplicationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
