using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/Racks_Get.json
// this example is just showing the usage of "Racks_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RackResource created on azure
// for more information of creating RackResource, please refer to the document of RackResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string rackName = "rackName";
ResourceIdentifier rackResourceId = RackResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, rackName);
RackResource rack = client.GetRackResource(rackResourceId);

// invoke the operation
RackResource result = await rack.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RackData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
