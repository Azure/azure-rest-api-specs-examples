using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CustomerInsights;
using Azure.ResourceManager.CustomerInsights.Models;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/LinksDelete.json
// this example is just showing the usage of "Links_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LinkResourceFormatResource created on azure
// for more information of creating LinkResourceFormatResource, please refer to the document of LinkResourceFormatResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
string linkName = "linkTest4806";
ResourceIdentifier linkResourceFormatResourceId = LinkResourceFormatResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName, linkName);
LinkResourceFormatResource linkResourceFormat = client.GetLinkResourceFormatResource(linkResourceFormatResourceId);

// invoke the operation
await linkResourceFormat.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
