using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Purview.Models;
using Azure.ResourceManager.Purview;

// Generated from example definition: specification/purview/resource-manager/Microsoft.Purview/preview/2023-05-01-preview/examples/PrivateEndpointConnections_Get.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PurviewAccountResource created on azure
// for more information of creating PurviewAccountResource, please refer to the document of PurviewAccountResource
string subscriptionId = "12345678-1234-1234-12345678abc";
string resourceGroupName = "SampleResourceGroup";
string accountName = "account1";
ResourceIdentifier purviewAccountResourceId = PurviewAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
PurviewAccountResource purviewAccount = client.GetPurviewAccountResource(purviewAccountResourceId);

// get the collection of this PurviewPrivateEndpointConnectionResource
PurviewPrivateEndpointConnectionCollection collection = purviewAccount.GetPurviewPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "privateEndpointConnection1";
NullableResponse<PurviewPrivateEndpointConnectionResource> response = await collection.GetIfExistsAsync(privateEndpointConnectionName);
PurviewPrivateEndpointConnectionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PurviewPrivateEndpointConnectionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
