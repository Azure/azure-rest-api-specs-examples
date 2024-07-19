using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-05-15-preview/examples/CosmosDBGraphResourceGet.json
// this example is just showing the usage of "GraphResources_GetGraph" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GraphResourceGetResultResource created on azure
// for more information of creating GraphResourceGetResultResource, please refer to the document of GraphResourceGetResultResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string graphName = "graphName";
ResourceIdentifier graphResourceGetResultResourceId = GraphResourceGetResultResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, graphName);
GraphResourceGetResultResource graphResourceGetResult = client.GetGraphResourceGetResultResource(graphResourceGetResultResourceId);

// invoke the operation
GraphResourceGetResultResource result = await graphResourceGetResult.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GraphResourceGetResultData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
