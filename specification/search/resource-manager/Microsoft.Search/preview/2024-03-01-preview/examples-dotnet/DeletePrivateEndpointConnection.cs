using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Search.Models;
using Azure.ResourceManager.Search;

// Generated from example definition: specification/search/resource-manager/Microsoft.Search/preview/2024-03-01-preview/examples/DeletePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SearchPrivateEndpointConnectionResource created on azure
// for more information of creating SearchPrivateEndpointConnectionResource, please refer to the document of SearchPrivateEndpointConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string searchServiceName = "mysearchservice";
string privateEndpointConnectionName = "testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546";
ResourceIdentifier searchPrivateEndpointConnectionResourceId = SearchPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, searchServiceName, privateEndpointConnectionName);
SearchPrivateEndpointConnectionResource searchPrivateEndpointConnection = client.GetSearchPrivateEndpointConnectionResource(searchPrivateEndpointConnectionResourceId);

// invoke the operation
ArmOperation<SearchPrivateEndpointConnectionResource> lro = await searchPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);
SearchPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SearchPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
