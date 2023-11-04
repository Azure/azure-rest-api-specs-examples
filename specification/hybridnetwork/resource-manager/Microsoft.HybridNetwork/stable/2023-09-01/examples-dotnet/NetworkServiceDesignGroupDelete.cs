using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkServiceDesignGroupDelete.json
// this example is just showing the usage of "NetworkServiceDesignGroups_Delete" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
await networkServiceDesignGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
