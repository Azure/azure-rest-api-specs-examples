using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationNetworks_ListByReplicationFabrics.json
// this example is just showing the usage of "ReplicationNetworks_ListByReplicationFabrics" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryFabricResource created on azure
// for more information of creating SiteRecoveryFabricResource, please refer to the document of SiteRecoveryFabricResource
string subscriptionId = "9112a37f-0f3e-46ec-9c00-060c6edca071";
string resourceGroupName = "srcBvte2a14C27";
string resourceName = "srce2avaultbvtaC27";
string fabricName = "b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac";
ResourceIdentifier siteRecoveryFabricResourceId = SiteRecoveryFabricResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName);
SiteRecoveryFabricResource siteRecoveryFabric = client.GetSiteRecoveryFabricResource(siteRecoveryFabricResourceId);

// get the collection of this SiteRecoveryNetworkResource
SiteRecoveryNetworkCollection collection = siteRecoveryFabric.GetSiteRecoveryNetworks();

// invoke the operation and iterate over the result
await foreach (SiteRecoveryNetworkResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SiteRecoveryNetworkData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
