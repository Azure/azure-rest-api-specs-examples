using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/bookmarks/GetBookmarkById.json
// this example is just showing the usage of "Bookmarks_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsBookmarkResource created on azure
// for more information of creating SecurityInsightsBookmarkResource, please refer to the document of SecurityInsightsBookmarkResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string bookmarkId = "73e01a99-5cd7-4139-a149-9f2736ff2ab5";
ResourceIdentifier securityInsightsBookmarkResourceId = SecurityInsightsBookmarkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, bookmarkId);
SecurityInsightsBookmarkResource securityInsightsBookmark = client.GetSecurityInsightsBookmarkResource(securityInsightsBookmarkResourceId);

// invoke the operation
SecurityInsightsBookmarkResource result = await securityInsightsBookmark.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsBookmarkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
