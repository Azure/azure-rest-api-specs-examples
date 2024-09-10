using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/bookmarks/relations/CreateBookmarkRelation.json
// this example is just showing the usage of "BookmarkRelations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsBookmarkResource created on azure
// for more information of creating SecurityInsightsBookmarkResource, please refer to the document of SecurityInsightsBookmarkResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string bookmarkId = "2216d0e1-91e3-4902-89fd-d2df8c535096";
ResourceIdentifier securityInsightsBookmarkResourceId = SecurityInsightsBookmarkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, bookmarkId);
SecurityInsightsBookmarkResource securityInsightsBookmark = client.GetSecurityInsightsBookmarkResource(securityInsightsBookmarkResourceId);

// get the collection of this SecurityInsightsBookmarkRelationResource
SecurityInsightsBookmarkRelationCollection collection = securityInsightsBookmark.GetSecurityInsightsBookmarkRelations();

// invoke the operation
string relationName = "4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014";
SecurityInsightsIncidentRelationData data = new SecurityInsightsIncidentRelationData()
{
    RelatedResourceId = new ResourceIdentifier("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/afbd324f-6c48-459c-8710-8d1e1cd03812"),
};
ArmOperation<SecurityInsightsBookmarkRelationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, relationName, data);
SecurityInsightsBookmarkRelationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsIncidentRelationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
