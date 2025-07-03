using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Avs;

// Generated from example definition: 2024-09-01/Hosts_Get.json
// this example is just showing the usage of "Host_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AvsPrivateCloudClusterResource created on azure
// for more information of creating AvsPrivateCloudClusterResource, please refer to the document of AvsPrivateCloudClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string clusterName = "cluster1";
ResourceIdentifier avsPrivateCloudClusterResourceId = AvsPrivateCloudClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, clusterName);
AvsPrivateCloudClusterResource avsPrivateCloudCluster = client.GetAvsPrivateCloudClusterResource(avsPrivateCloudClusterResourceId);

// get the collection of this AvsHostResource
AvsHostCollection collection = avsPrivateCloudCluster.GetAvsHosts();

// invoke the operation
string hostId = "esx03-r52.1111111111111111111.westcentralus.prod.azure.com";
NullableResponse<AvsHostResource> response = await collection.GetIfExistsAsync(hostId);
AvsHostResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AvsHostData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
