using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HardwareSecurityModules.Models;
using Azure.ResourceManager.HardwareSecurityModules;

// Generated from example definition: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-preview/examples/CloudHsmClusterPrivateEndpointConnection_Create_MaximumSet_Gen.json
// this example is just showing the usage of "CloudHsmClusterPrivateEndpointConnections_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CloudHsmClusterResource created on azure
// for more information of creating CloudHsmClusterResource, please refer to the document of CloudHsmClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgcloudhsm";
string cloudHsmClusterName = "chsm1";
ResourceIdentifier cloudHsmClusterResourceId = CloudHsmClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudHsmClusterName);
CloudHsmClusterResource cloudHsmCluster = client.GetCloudHsmClusterResource(cloudHsmClusterResourceId);

// get the collection of this CloudHsmClusterPrivateEndpointConnectionResource
CloudHsmClusterPrivateEndpointConnectionCollection collection = cloudHsmCluster.GetCloudHsmClusterPrivateEndpointConnections();

// invoke the operation
string peConnectionName = "sample-pec";
CloudHsmClusterPrivateEndpointConnectionData data = new CloudHsmClusterPrivateEndpointConnectionData()
{
    Properties = new CloudHsmClusterPrivateEndpointConnectionProperties(new CloudHsmClusterPrivateLinkServiceConnectionState()
    {
        Status = CloudHsmClusterPrivateEndpointServiceConnectionStatus.Approved,
        Description = "My name is Joe and I'm approving this.",
    }),
};
ArmOperation<CloudHsmClusterPrivateEndpointConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, peConnectionName, data);
CloudHsmClusterPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CloudHsmClusterPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
