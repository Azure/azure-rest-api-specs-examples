using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HardwareSecurityModules;
using Azure.ResourceManager.HardwareSecurityModules.Models;

// Generated from example definition: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2023-12-10-preview/examples/CloudHsmClusterPrivateEndpointConnection_Get_MaximumSet_Gen.json
// this example is just showing the usage of "CloudHsmClusterPrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this HardwareSecurityModulesPrivateEndpointConnectionResource
HardwareSecurityModulesPrivateEndpointConnectionCollection collection = cloudHsmCluster.GetHardwareSecurityModulesPrivateEndpointConnections();

// invoke the operation
string peConnectionName = "sample-pec";
NullableResponse<HardwareSecurityModulesPrivateEndpointConnectionResource> response = await collection.GetIfExistsAsync(peConnectionName);
HardwareSecurityModulesPrivateEndpointConnectionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HardwareSecurityModulesPrivateEndpointConnectionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
