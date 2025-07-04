using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/CloudServicesNetworks_Patch.json
// this example is just showing the usage of "CloudServicesNetworks_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudCloudServicesNetworkResource created on azure
// for more information of creating NetworkCloudCloudServicesNetworkResource, please refer to the document of NetworkCloudCloudServicesNetworkResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string cloudServicesNetworkName = "cloudServicesNetworkName";
ResourceIdentifier networkCloudCloudServicesNetworkResourceId = NetworkCloudCloudServicesNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudServicesNetworkName);
NetworkCloudCloudServicesNetworkResource networkCloudCloudServicesNetwork = client.GetNetworkCloudCloudServicesNetworkResource(networkCloudCloudServicesNetworkResourceId);

// invoke the operation
NetworkCloudCloudServicesNetworkPatch patch = new NetworkCloudCloudServicesNetworkPatch
{
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2"
    },
    AdditionalEgressEndpoints = {new EgressEndpoint("azure-resource-management", new EndpointDependency[]
    {
    new EndpointDependency("storageaccountex.blob.core.windows.net")
    {
    Port = 443L,
    }
    })},
    EnableDefaultEgressEndpoints = CloudServicesNetworkEnableDefaultEgressEndpoint.False,
};
ArmOperation<NetworkCloudCloudServicesNetworkResource> lro = await networkCloudCloudServicesNetwork.UpdateAsync(WaitUntil.Completed, patch);
NetworkCloudCloudServicesNetworkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkCloudCloudServicesNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
