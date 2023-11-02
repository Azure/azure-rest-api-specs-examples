using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceNetworking;
using Azure.ResourceManager.ServiceNetworking.Models;

// Generated from example definition: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/stable/2023-11-01/examples/FrontendDelete.json
// this example is just showing the usage of "FrontendsInterface_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontendResource created on azure
// for more information of creating FrontendResource, please refer to the document of FrontendResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string trafficControllerName = "tc1";
string frontendName = "fe1";
ResourceIdentifier frontendResourceId = FrontendResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, trafficControllerName, frontendName);
FrontendResource frontend = client.GetFrontendResource(frontendResourceId);

// invoke the operation
await frontend.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
