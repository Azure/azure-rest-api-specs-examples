using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/FileSharesLease_Break.json
// this example is just showing the usage of "FileShares_Lease" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FileShareResource created on azure
// for more information of creating FileShareResource, please refer to the document of FileShareResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res3376";
string accountName = "sto328";
string shareName = "share12";
ResourceIdentifier fileShareResourceId = FileShareResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, shareName);
FileShareResource fileShare = client.GetFileShareResource(fileShareResourceId);

// invoke the operation
LeaseShareContent content = new LeaseShareContent(LeaseShareAction.Break)
{
    LeaseId = "8698f513-fa75-44a1-b8eb-30ba336af27d",
    BreakPeriod = null,
    LeaseDuration = null,
    ProposedLeaseId = null,
};
LeaseShareResponse result = await fileShare.LeaseAsync(content: content);

Console.WriteLine($"Succeeded: {result}");
