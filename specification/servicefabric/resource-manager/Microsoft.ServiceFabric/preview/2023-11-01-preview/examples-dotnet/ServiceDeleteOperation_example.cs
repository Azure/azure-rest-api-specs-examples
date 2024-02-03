using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceFabric;
using Azure.ResourceManager.ServiceFabric.Models;

// Generated from example definition: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/preview/2023-11-01-preview/examples/ServiceDeleteOperation_example.json
// this example is just showing the usage of "Services_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceFabricServiceResource created on azure
// for more information of creating ServiceFabricServiceResource, please refer to the document of ServiceFabricServiceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string clusterName = "myCluster";
string applicationName = "myApp";
string serviceName = "myService";
ResourceIdentifier serviceFabricServiceResourceId = ServiceFabricServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, applicationName, serviceName);
ServiceFabricServiceResource serviceFabricService = client.GetServiceFabricServiceResource(serviceFabricServiceResourceId);

// invoke the operation
await serviceFabricService.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
