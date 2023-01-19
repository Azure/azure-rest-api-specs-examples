using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CustomerInsights;
using Azure.ResourceManager.CustomerInsights.Models;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ProfileGetEnrichingKpis.json
// this example is just showing the usage of "Profiles_GetEnrichingKpis" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProfileResourceFormatResource created on azure
// for more information of creating ProfileResourceFormatResource, please refer to the document of ProfileResourceFormatResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
string profileName = "TestProfileType396";
ResourceIdentifier profileResourceFormatResourceId = ProfileResourceFormatResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName, profileName);
ProfileResourceFormatResource profileResourceFormat = client.GetProfileResourceFormatResource(profileResourceFormatResourceId);

// invoke the operation and iterate over the result
await foreach (KpiDefinition item in profileResourceFormat.GetEnrichingKpisAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
