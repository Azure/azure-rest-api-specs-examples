using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/bookmarks/CreateBookmark.json
// this example is just showing the usage of "Bookmarks_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
SecurityInsightsBookmarkData data = new SecurityInsightsBookmarkData()
{
    CreatedOn = DateTimeOffset.Parse("2021-09-01T13:15:30Z"),
    CreatedBy = new SecurityInsightsUserInfo()
    {
        ObjectId = Guid.Parse("2046feea-040d-4a46-9e2b-91c2941bfa70"),
    },
    DisplayName = "My bookmark",
    Labels =
    {
    "Tag1","Tag2"
    },
    Notes = "Found a suspicious activity",
    Query = "SecurityEvent | where TimeGenerated > ago(1d) and TimeGenerated < ago(2d)",
    QueryResult = "Security Event query result",
    UpdatedOn = DateTimeOffset.Parse("2021-09-01T13:15:30Z"),
    UpdatedBy = new SecurityInsightsUserInfo()
    {
        ObjectId = Guid.Parse("2046feea-040d-4a46-9e2b-91c2941bfa70"),
    },
    EntityMappings =
    {
    new BookmarkEntityMappings()
    {
    EntityType = "Account",
    FieldMappings =
    {
    new EntityFieldMapping()
    {
    Identifier = "Fullname",
    Value = "johndoe@microsoft.com",
    }
    },
    }
    },
    Tactics =
    {
    SecurityInsightsAttackTactic.Execution
    },
    Techniques =
    {
    "T1609"
    },
    ETag = new ETag("\"0300bf09-0000-0000-0000-5c37296e0000\""),
};
ArmOperation<SecurityInsightsBookmarkResource> lro = await securityInsightsBookmark.UpdateAsync(WaitUntil.Completed, data);
SecurityInsightsBookmarkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsBookmarkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
