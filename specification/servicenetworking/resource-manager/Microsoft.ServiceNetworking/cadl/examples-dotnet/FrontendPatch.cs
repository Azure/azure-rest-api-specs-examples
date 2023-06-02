using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceNetworking;
using Azure.ResourceManager.ServiceNetworking.Models;

// Generated from example definition: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/FrontendPatch.json
// this example is just showing the usage of "FrontendsInterface_Update" operation, for the dependent resources, they will have to be created separately.

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
FrontendPatch patch = new FrontendPatch();
FrontendResource result = await frontend.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontendData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
