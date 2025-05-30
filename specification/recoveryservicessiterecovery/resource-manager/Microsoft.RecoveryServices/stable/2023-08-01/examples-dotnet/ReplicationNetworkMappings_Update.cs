using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationNetworkMappings_Update.json
// this example is just showing the usage of "ReplicationNetworkMappings_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryNetworkMappingResource created on azure
// for more information of creating SiteRecoveryNetworkMappingResource, please refer to the document of SiteRecoveryNetworkMappingResource
string subscriptionId = "9112a37f-0f3e-46ec-9c00-060c6edca071";
string resourceGroupName = "srcBvte2a14C27";
string resourceName = "srce2avaultbvtaC27";
string fabricName = "b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac";
string networkName = "e2267b5c-2650-49bd-ab3f-d66aae694c06";
string networkMappingName = "corpe2amap";
ResourceIdentifier siteRecoveryNetworkMappingResourceId = SiteRecoveryNetworkMappingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, networkName, networkMappingName);
SiteRecoveryNetworkMappingResource siteRecoveryNetworkMapping = client.GetSiteRecoveryNetworkMappingResource(siteRecoveryNetworkMappingResourceId);

// invoke the operation
SiteRecoveryNetworkMappingPatch patch = new SiteRecoveryNetworkMappingPatch
{
    Properties = new UpdateNetworkMappingProperties
    {
        RecoveryFabricName = "Microsoft Azure",
        RecoveryNetworkId = new ResourceIdentifier("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai2"),
        FabricSpecificDetails = new VmmToAzureUpdateNetworkMappingContent(),
    },
};
ArmOperation<SiteRecoveryNetworkMappingResource> lro = await siteRecoveryNetworkMapping.UpdateAsync(WaitUntil.Completed, patch);
SiteRecoveryNetworkMappingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryNetworkMappingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
