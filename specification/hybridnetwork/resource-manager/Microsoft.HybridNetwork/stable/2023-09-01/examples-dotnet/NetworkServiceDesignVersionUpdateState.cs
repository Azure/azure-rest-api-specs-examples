using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkServiceDesignVersionUpdateState.json
// this example is just showing the usage of "NetworkServiceDesignVersions_updateState" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkServiceDesignVersionResource created on azure
// for more information of creating NetworkServiceDesignVersionResource, please refer to the document of NetworkServiceDesignVersionResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string publisherName = "TestPublisher";
string networkServiceDesignGroupName = "TestNetworkServiceDesignGroupName";
string networkServiceDesignVersionName = "1.0.0";
ResourceIdentifier networkServiceDesignVersionResourceId = NetworkServiceDesignVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, publisherName, networkServiceDesignGroupName, networkServiceDesignVersionName);
NetworkServiceDesignVersionResource networkServiceDesignVersion = client.GetNetworkServiceDesignVersionResource(networkServiceDesignVersionResourceId);

// invoke the operation
NetworkServiceDesignVersionUpdateState networkServiceDesignVersionUpdateState = new NetworkServiceDesignVersionUpdateState()
{
    VersionState = VersionState.Active,
};
ArmOperation<NetworkServiceDesignVersionUpdateState> lro = await networkServiceDesignVersion.UpdateStateAsync(WaitUntil.Completed, networkServiceDesignVersionUpdateState);
NetworkServiceDesignVersionUpdateState result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
