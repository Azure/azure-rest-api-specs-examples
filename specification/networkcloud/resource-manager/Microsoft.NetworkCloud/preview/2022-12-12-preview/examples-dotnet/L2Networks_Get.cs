using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/L2Networks_Get.json
// this example is just showing the usage of "L2Networks_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this L2NetworkResource created on azure
// for more information of creating L2NetworkResource, please refer to the document of L2NetworkResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string l2NetworkName = "l2NetworkName";
ResourceIdentifier l2NetworkResourceId = L2NetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l2NetworkName);
L2NetworkResource l2Network = client.GetL2NetworkResource(l2NetworkResourceId);

// invoke the operation
L2NetworkResource result = await l2Network.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
L2NetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
