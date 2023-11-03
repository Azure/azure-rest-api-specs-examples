using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkServiceDesignVersionGet.json
// this example is just showing the usage of "NetworkServiceDesignVersions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkServiceDesignGroupResource created on azure
// for more information of creating NetworkServiceDesignGroupResource, please refer to the document of NetworkServiceDesignGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string publisherName = "TestPublisher";
string networkServiceDesignGroupName = "TestNetworkServiceDesignGroupName";
ResourceIdentifier networkServiceDesignGroupResourceId = NetworkServiceDesignGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, publisherName, networkServiceDesignGroupName);
NetworkServiceDesignGroupResource networkServiceDesignGroup = client.GetNetworkServiceDesignGroupResource(networkServiceDesignGroupResourceId);

// get the collection of this NetworkServiceDesignVersionResource
NetworkServiceDesignVersionCollection collection = networkServiceDesignGroup.GetNetworkServiceDesignVersions();

// invoke the operation
string networkServiceDesignVersionName = "1.0.0";
NullableResponse<NetworkServiceDesignVersionResource> response = await collection.GetIfExistsAsync(networkServiceDesignVersionName);
NetworkServiceDesignVersionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkServiceDesignVersionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
