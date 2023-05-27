using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/CloudServicesNetworks_Patch.json
// this example is just showing the usage of "CloudServicesNetworks_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CloudServicesNetworkResource created on azure
// for more information of creating CloudServicesNetworkResource, please refer to the document of CloudServicesNetworkResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string cloudServicesNetworkName = "cloudServicesNetworkName";
ResourceIdentifier cloudServicesNetworkResourceId = CloudServicesNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudServicesNetworkName);
CloudServicesNetworkResource cloudServicesNetwork = client.GetCloudServicesNetworkResource(cloudServicesNetworkResourceId);

// invoke the operation
CloudServicesNetworkPatch patch = new CloudServicesNetworkPatch()
{
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2",
    },
    AdditionalEgressEndpoints =
    {
    new EgressEndpoint("azure-resource-management",new EndpointDependency[]
    {
    new EndpointDependency("https://storageaccountex.blob.core.windows.net")
    {
    Port = 443,
    }
    })
    },
    EnableDefaultEgressEndpoints = CloudServicesNetworkEnableDefaultEgressEndpoint.False,
};
ArmOperation<CloudServicesNetworkResource> lro = await cloudServicesNetwork.UpdateAsync(WaitUntil.Completed, patch);
CloudServicesNetworkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CloudServicesNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
