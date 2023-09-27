using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationNetworkMappings_ListByReplicationNetworks.json
// this example is just showing the usage of "ReplicationNetworkMappings_ListByReplicationNetworks" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryNetworkResource created on azure
// for more information of creating SiteRecoveryNetworkResource, please refer to the document of SiteRecoveryNetworkResource
string subscriptionId = "9112a37f-0f3e-46ec-9c00-060c6edca071";
string resourceGroupName = "srcBvte2a14C27";
string resourceName = "srce2avaultbvtaC27";
string fabricName = "b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac";
string networkName = "e2267b5c-2650-49bd-ab3f-d66aae694c06";
ResourceIdentifier siteRecoveryNetworkResourceId = SiteRecoveryNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, networkName);
SiteRecoveryNetworkResource siteRecoveryNetwork = client.GetSiteRecoveryNetworkResource(siteRecoveryNetworkResourceId);

// get the collection of this SiteRecoveryNetworkMappingResource
SiteRecoveryNetworkMappingCollection collection = siteRecoveryNetwork.GetSiteRecoveryNetworkMappings();

// invoke the operation and iterate over the result
await foreach (SiteRecoveryNetworkMappingResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SiteRecoveryNetworkMappingData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
