using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/watchlists/CreateWatchlist.json
// this example is just showing the usage of "Watchlists_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OperationalInsightsWorkspaceSecurityInsightsResource created on azure
// for more information of creating OperationalInsightsWorkspaceSecurityInsightsResource, please refer to the document of OperationalInsightsWorkspaceSecurityInsightsResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
ResourceIdentifier operationalInsightsWorkspaceSecurityInsightsResourceId = OperationalInsightsWorkspaceSecurityInsightsResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
OperationalInsightsWorkspaceSecurityInsightsResource operationalInsightsWorkspaceSecurityInsights = client.GetOperationalInsightsWorkspaceSecurityInsightsResource(operationalInsightsWorkspaceSecurityInsightsResourceId);

// get the collection of this SecurityInsightsWatchlistResource
SecurityInsightsWatchlistCollection collection = operationalInsightsWorkspaceSecurityInsights.GetSecurityInsightsWatchlists();

// invoke the operation
string watchlistAlias = "highValueAsset";
SecurityInsightsWatchlistData data = new SecurityInsightsWatchlistData()
{
    DisplayName = "High Value Assets Watchlist",
    Provider = "Microsoft",
    SourceString = "watchlist.csv",
    SourceType = WatchlistSourceType.LocalFile,
    Description = "Watchlist from CSV content",
    ItemsSearchKey = "header1",
    ETag = new ETag("\"0300bf09-0000-0000-0000-5c37296e0000\""),
};
ArmOperation<SecurityInsightsWatchlistResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, watchlistAlias, data);
SecurityInsightsWatchlistResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsWatchlistData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
