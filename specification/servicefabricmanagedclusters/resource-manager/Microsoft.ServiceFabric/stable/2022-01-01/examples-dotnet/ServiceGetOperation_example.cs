using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceFabricManagedClusters;
using Azure.ResourceManager.ServiceFabricManagedClusters.Models;

// Generated from example definition: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2022-01-01/examples/ServiceGetOperation_example.json
// this example is just showing the usage of "Services_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ServiceFabricManagedServiceResource
ServiceFabricManagedServiceCollection collection = serviceFabricManagedApplication.GetServiceFabricManagedServices();

// invoke the operation
string serviceName = "myService";
bool result = await collection.ExistsAsync(serviceName);

Console.WriteLine($"Succeeded: {result}");
