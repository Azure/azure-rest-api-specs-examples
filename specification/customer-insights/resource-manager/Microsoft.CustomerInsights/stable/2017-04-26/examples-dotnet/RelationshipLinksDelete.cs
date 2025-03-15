using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CustomerInsights.Models;
using Azure.ResourceManager.CustomerInsights;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipLinksDelete.json
// this example is just showing the usage of "RelationshipLinks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RelationshipLinkResourceFormatResource created on azure
// for more information of creating RelationshipLinkResourceFormatResource, please refer to the document of RelationshipLinkResourceFormatResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
string relationshipLinkName = "Somelink";
ResourceIdentifier relationshipLinkResourceFormatResourceId = RelationshipLinkResourceFormatResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName, relationshipLinkName);
RelationshipLinkResourceFormatResource relationshipLinkResourceFormat = client.GetRelationshipLinkResourceFormatResource(relationshipLinkResourceFormatResourceId);

// invoke the operation
await relationshipLinkResourceFormat.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
