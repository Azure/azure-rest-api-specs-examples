using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/GetPrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CognitiveServicesAccountResource created on azure
// for more information of creating CognitiveServicesAccountResource, please refer to the document of CognitiveServicesAccountResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "res6977";
string accountName = "sto2527";
ResourceIdentifier cognitiveServicesAccountResourceId = CognitiveServicesAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
CognitiveServicesAccountResource cognitiveServicesAccount = client.GetCognitiveServicesAccountResource(cognitiveServicesAccountResourceId);

// get the collection of this CognitiveServicesPrivateEndpointConnectionResource
CognitiveServicesPrivateEndpointConnectionCollection collection = cognitiveServicesAccount.GetCognitiveServicesPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "{privateEndpointConnectionName}";
NullableResponse<CognitiveServicesPrivateEndpointConnectionResource> response = await collection.GetIfExistsAsync(privateEndpointConnectionName);
CognitiveServicesPrivateEndpointConnectionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CognitiveServicesPrivateEndpointConnectionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
