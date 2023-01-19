using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CustomerInsights;
using Azure.ResourceManager.CustomerInsights.Models;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/InteractionsSuggestRelationshipLinks.json
// this example is just showing the usage of "Interactions_SuggestRelationshipLinks" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InteractionResourceFormatResource created on azure
// for more information of creating InteractionResourceFormatResource, please refer to the document of InteractionResourceFormatResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
string interactionName = "Deposit";
ResourceIdentifier interactionResourceFormatResourceId = InteractionResourceFormatResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName, interactionName);
InteractionResourceFormatResource interactionResourceFormat = client.GetInteractionResourceFormatResource(interactionResourceFormatResourceId);

// invoke the operation
SuggestRelationshipLinksResponse result = await interactionResourceFormat.SuggestRelationshipLinksAsync();

Console.WriteLine($"Succeeded: {result}");
