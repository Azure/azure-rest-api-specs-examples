using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CustomerInsights;
using Azure.ResourceManager.CustomerInsights.Models;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipsGet.json
// this example is just showing the usage of "Relationships_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HubResource created on azure
// for more information of creating HubResource, please refer to the document of HubResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
ResourceIdentifier hubResourceId = HubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName);
HubResource hub = client.GetHubResource(hubResourceId);

// get the collection of this RelationshipResourceFormatResource
RelationshipResourceFormatCollection collection = hub.GetRelationshipResourceFormats();

// invoke the operation
string relationshipName = "SomeRelationship";
bool result = await collection.ExistsAsync(relationshipName);

Console.WriteLine($"Succeeded: {result}");
