using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CustomerInsights;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ViewsDelete.json
// this example is just showing the usage of "Views_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ViewResourceFormatResource created on azure
// for more information of creating ViewResourceFormatResource, please refer to the document of ViewResourceFormatResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
string viewName = "testView";
ResourceIdentifier viewResourceFormatResourceId = ViewResourceFormatResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName, viewName);
ViewResourceFormatResource viewResourceFormat = client.GetViewResourceFormatResource(viewResourceFormatResourceId);

// invoke the operation
string userId = "*";
await viewResourceFormat.DeleteAsync(WaitUntil.Completed, userId);

Console.WriteLine($"Succeeded");
