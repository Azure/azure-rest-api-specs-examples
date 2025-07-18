using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/L3Networks_Get.json
// this example is just showing the usage of "L3Networks_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudL3NetworkResource created on azure
// for more information of creating NetworkCloudL3NetworkResource, please refer to the document of NetworkCloudL3NetworkResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string l3NetworkName = "l3NetworkName";
ResourceIdentifier networkCloudL3NetworkResourceId = NetworkCloudL3NetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l3NetworkName);
NetworkCloudL3NetworkResource networkCloudL3Network = client.GetNetworkCloudL3NetworkResource(networkCloudL3NetworkResourceId);

// invoke the operation
NetworkCloudL3NetworkResource result = await networkCloudL3Network.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkCloudL3NetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
