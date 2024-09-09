using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/watchlists/CreateWatchlistItem.json
// this example is just showing the usage of "WatchlistItems_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsWatchlistResource created on azure
// for more information of creating SecurityInsightsWatchlistResource, please refer to the document of SecurityInsightsWatchlistResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string watchlistAlias = "highValueAsset";
ResourceIdentifier securityInsightsWatchlistResourceId = SecurityInsightsWatchlistResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, watchlistAlias);
SecurityInsightsWatchlistResource securityInsightsWatchlist = client.GetSecurityInsightsWatchlistResource(securityInsightsWatchlistResourceId);

// get the collection of this SecurityInsightsWatchlistItemResource
SecurityInsightsWatchlistItemCollection collection = securityInsightsWatchlist.GetSecurityInsightsWatchlistItems();

// invoke the operation
string watchlistItemId = "82ba292c-dc97-4dfc-969d-d4dd9e666842";
SecurityInsightsWatchlistItemData data = new SecurityInsightsWatchlistItemData()
{
    ItemsKeyValueDictionary =
    {
    ["Business tier"] = BinaryData.FromString("\"10.0.2.0/24\""),
    ["Data tier"] = BinaryData.FromString("\"10.0.2.0/24\""),
    ["Gateway subnet"] = BinaryData.FromString("\"10.0.255.224/27\""),
    ["Private DMZ in"] = BinaryData.FromString("\"10.0.0.0/27\""),
    ["Public DMZ out"] = BinaryData.FromString("\"10.0.0.96/27\""),
    ["Web Tier"] = BinaryData.FromString("\"10.0.1.0/24\""),
    },
    ETag = new ETag("0300bf09-0000-0000-0000-5c37296e0000"),
};
ArmOperation<SecurityInsightsWatchlistItemResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, watchlistItemId, data);
SecurityInsightsWatchlistItemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsWatchlistItemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
