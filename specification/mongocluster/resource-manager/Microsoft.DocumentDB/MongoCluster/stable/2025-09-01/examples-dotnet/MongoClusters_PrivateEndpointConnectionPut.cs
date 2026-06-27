using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MongoCluster.Models;
using Azure.ResourceManager.MongoCluster;

// Generated from example definition: 2025-09-01/MongoClusters_PrivateEndpointConnectionPut.json
// this example is just showing the usage of "PrivateEndpointConnectionResource_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MongoClusterPrivateEndpointConnectionResource created on azure
// for more information of creating MongoClusterPrivateEndpointConnectionResource, please refer to the document of MongoClusterPrivateEndpointConnectionResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string mongoClusterName = "myMongoCluster";
string privateEndpointConnectionName = "pecTest";
ResourceIdentifier mongoClusterPrivateEndpointConnectionResourceId = MongoClusterPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, mongoClusterName, privateEndpointConnectionName);
MongoClusterPrivateEndpointConnectionResource mongoClusterPrivateEndpointConnectionResource = client.GetMongoClusterPrivateEndpointConnectionResource(mongoClusterPrivateEndpointConnectionResourceId);

// invoke the operation
MongoClusterPrivateEndpointConnectionResourceData data = new MongoClusterPrivateEndpointConnectionResourceData
{
    Properties = new MongoClusterPrivateEndpointConnectionProperties(new MongoClusterPrivateLinkServiceConnectionState
    {
        Status = MongoClusterPrivateEndpointServiceConnectionStatus.Approved,
        Description = "Auto-Approved",
    }),
};
ArmOperation<MongoClusterPrivateEndpointConnectionResource> lro = await mongoClusterPrivateEndpointConnectionResource.UpdateAsync(WaitUntil.Completed, data);
MongoClusterPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MongoClusterPrivateEndpointConnectionResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
